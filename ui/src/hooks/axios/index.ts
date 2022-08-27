import axios from "axios";
import type {
  AxiosError,
  AxiosRequestConfig,
  AxiosResponse,
  CancelTokenSource,
} from "axios";
import { ref, shallowRef } from "vue";
import { debounce } from "howtools";
export * from "./help";

export type RequestInterceptor = (
  config: AxiosRequestConfig
) => AxiosRequestConfig;
export type ResponseInterceptor = (response: AxiosResponse) => AxiosResponse;
export type ErrResponseInterceptor = (err: any) => any;

export interface HowAxiosRequestConfig extends AxiosRequestConfig {
  path?: any;
}

export interface HowVesAxiosOptions {
  instanceConfig: AxiosRequestConfig;
  requestInterceptor?: RequestInterceptor;
  responseInterceptor?: ResponseInterceptor;
  errResponseInterceptor?: ErrResponseInterceptor;
}

export interface HowVesExRequestOptions {
  immediate?: boolean;
  delay?: number;
  isDebounce?: boolean;
  defaultVal?: any;
}

export function createAxios(options: HowVesAxiosOptions) {
  const {
    instanceConfig: config,
    requestInterceptor,
    responseInterceptor,
    errResponseInterceptor,
  } = options;

  const server = axios.create({
    ...config,
    baseURL: config.baseURL,
  });

  // 请求拦截器
  server.interceptors.request.use((config) => {
    const c = config as HowAxiosRequestConfig;
    for (const key in c.path) {
      config.url = config.url?.replace(`{${key}}`, c.path[key]);
    }
    delete c.path;
    if (!requestInterceptor) return c;
    return requestInterceptor(c as AxiosRequestConfig);
  });

  // 响应拦截器
  server.interceptors.response.use(
    (response) => {
      // 设置不允许修改原始data
      responseInterceptor && responseInterceptor(response);
      return response;
    },
    (err: any) => {
      // 失败拦截处理
      errResponseInterceptor && errResponseInterceptor(err);
      // return err
    }
  );

  // Axios hook
  /**
   * @param  {AxiosRequestConfig} config
   * @param  {HowVesExRequestOptions} options?
   * @returns
   */
  function useAxiosRequest<T>(
    config: HowAxiosRequestConfig,
    options?: HowVesExRequestOptions
  ) {
    const { isDebounce = true, defaultVal = {} } = options || {};

    let lastConf = config; //最后一次发出请求的配置
    const isLoading = shallowRef(false);
    const isFinished = shallowRef(false);
    const aborted = shallowRef(false); // 请求被中断
    const cancelToken: CancelTokenSource = axios.CancelToken.source();

    const loading = (loading: boolean) => {
      isLoading.value = loading;
      isFinished.value = !loading;
    };

    const abort = (message?: string) => {
      if (isFinished.value || !isLoading.value) return;

      cancelToken.cancel(message);
      aborted.value = true;
      isLoading.value = false;
      isFinished.value = false;
    };

    const response = ref<AxiosResponse<T>>(); //axios响应
    const data = ref<T>(defaultVal); //响应数据
    const error = ref<AxiosError<T>>(); // axios 错误响应
    const edata = ref<T>(); // axios 错误响应数据

    // 不是节流的方式
    const preRequest = ({
      params: p,
      data: d,
      path: pv,
    }: HowAxiosRequestConfig) => {
      const c = { ...config, params: p, data: d, path: pv };
      server
        .request({ ...c, cancelToken: cancelToken.token })
        .then((r) => {
          response.value = r;
          data.value = r.data;
          loading(false);
        })
        .catch((e: AxiosError) => {
          error.value = e as any;
          edata.value = e.response ? e.response.data : ("" as any);
          loading(false);
        });
    };

    const request = debounce(preRequest, (options && options.delay) || 1);

    const execute = (
      config: Pick<HowAxiosRequestConfig, "params" | "data" | "path"> = {
        params: {},
        data: {},
        path: {},
      }
    ): Promise<T> => {
      lastConf = config;

      loading(true);

      if (isDebounce) {
        request(config);
      } else {
        preRequest(config);
      }
      return new Promise((resolve) => {
        const resultInterval = setInterval(() => {
          if (isFinished.value) {
            clearInterval(resultInterval);
            resolve(data.value as T);
          }
        }, 100);
      });
    };

    // 立即执行
    if (options?.immediate)
      execute({
        path: config.path,
        params: config.params,
        data: config.data,
      });

    // 重新加载上次请求
    function reload() {
      execute(lastConf);
    }

    return {
      response,
      data,
      error,
      edata,
      execute,
      reload,
      aborted,
      abort,
      finished: isFinished,
      loading: isLoading,
    };
  }

  return {
    server,
    useAxiosRequest,
  };
}
