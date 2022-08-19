import { useDefRequest } from "@/utils/requests";

// 添加识别码
export function useCodeAdd() {
  return useDefRequest({
    url: "/api/code/add",
    method: "POST",
  });
}

// 获取code
export function useCodeGet() {
  return useDefRequest({
    url: "/api/code/get",
    method: "GET",
  });
}
