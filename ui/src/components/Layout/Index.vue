<template>
  <a-layout class="layout">
    <a-layout-header>
      <div class="logo" />
      <div style="font-size: 18px; color: aliceblue">
        {{ data.title }}
      </div>
    </a-layout-header>
    <a-layout-content>
      <div
        :style="{
          background: '#fff',
          padding: '24px',
          minHeight: 'calc(100vh - 240px)',
          maxWidth: '1200px',
          margin: '0 auto',
        }"
      >
        <router-view />
      </div>
    </a-layout-content>
    <a-layout-footer style="text-align: center">
      <img width="200" :src="data.qrcodeLink" />
      <div>{{ data.description }}</div>
    </a-layout-footer>
  </a-layout>
</template>
<script setup lang="ts">
import { useConfig } from "@/api/config";

const { data } = useConfig();
</script>
<style>
.site-layout-content {
  min-height: 280px;
  padding: 24px;
  background: #fff;
}
#components-layout-demo-top .logo {
  float: left;
  width: 120px;
  height: 31px;
  margin: 16px 24px 16px 0;
  background: rgba(255, 255, 255, 0.3);
}
.ant-row-rtl #components-layout-demo-top .logo {
  float: right;
  margin: 16px 0 16px 24px;
}

[data-theme="dark"] .site-layout-content {
  background: #141414;
}
</style>
