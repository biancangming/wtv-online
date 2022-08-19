import { useDefRequest } from "@/utils/requests";

export function useShareUpdate() {
  return useDefRequest({
    url: "/api/share/update",
    method: "POST",
  });
}

export function useShareUpdateStatus() {
  return useDefRequest({
    url: "/api/share/updateStatus",
    method: "POST",
  });
}

export function useShareGet() {
  return useDefRequest({
    url: "/api/share/get",
    method: "GET",
  });
}

// useStatus 使用状态 0 启用中 1 是废弃
export function useShareUrls(immediate = false, useStatus: 0 | 1) {
  return useDefRequest(
    {
      url: "/api/share/urls",
      params: {
        useStatus,
      },
    },
    {
      immediate,
      defaultVal: [],
    }
  );
}
