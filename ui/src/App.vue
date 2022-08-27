<template>
  <RouterView v-if="isRequired" />
  <Login v-else />
</template>
<script lang="ts" setup>
import { useIsRequiredLogin } from "./api/config";
import Login from "@/components/Login.vue";
import { computed } from "vue";
const { data } = useIsRequiredLogin();
const isRequired = computed(() => {
  const logined = JSON.parse(sessionStorage.getItem("logined") || "false");
  // 没有设置账号密码或者必须是登录状态
  return data.value === 0 || logined;
});
</script>
<style lang="less" scoped></style>
