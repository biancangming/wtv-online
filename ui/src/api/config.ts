import { useAxiosRequest } from "@/utils/requests";

export function useIsRequiredLogin() {
  return useAxiosRequest(
    {
      url: "/api/config/isRequiredLogin",
    },
    {
      immediate: true,
    }
  );
}

export function useConfig() {
  return useAxiosRequest<{
    title: string;
    qrcodeLink: string;
    description: string;
  }>(
    {
      url: "/api/config/config",
    },
    {
      immediate: true,
      defaultVal: {},
    }
  );
}

export function UseLogin() {
  return useAxiosRequest<string>({
    url: "/api/config/login",
    method: "POST",
  });
}
