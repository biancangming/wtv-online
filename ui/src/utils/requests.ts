import {
  createAxios,
  type HowAxiosRequestConfig,
  type HowVesExRequestOptions,
} from "@/hooks/axios";
import { message } from "ant-design-vue";
import { watchEffect, type Ref } from "vue";
import { watch } from "vue";

const { useAxiosRequest } = createAxios({
  instanceConfig: {
    baseURL: import.meta.env.VITE_LIVE_BASE as string,
  },
  requestInterceptor: (config) => {
    if (config.headers) {
      config.headers["Authorization"] = sessionStorage.getItem("token") || "";
    }
    return config;
  },
  errResponseInterceptor: (err) => {
    if (err.request.status == 401) {
      location.href = "/#/login";
    }
  },
});

export interface Result<T> {
  code: 0 | 1;
  msg: string;
  data: T;
}

// 默认请求列表
export function useDefRequest<T>(
  config: HowAxiosRequestConfig,
  options?: HowVesExRequestOptions
) {
  return useAxiosRequest<Result<T>>(config, options);
}

// 错误提示辅助函数
export function useDataTip(result: Ref<Result<any>> | any, error?: Ref<any>) {
  watch(result, () => {
    if (result.value.code == 0) {
      message.success(result.value.msg);
    } else {
      message.error(result.value.msg);
    }
  });

  if (error) {
    const msg = error.value?.response?.data.msg;
    watchEffect(() => msg && message.error(msg));
  }
}
