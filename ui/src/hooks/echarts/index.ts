import * as echarts from "echarts";

import type { ECharts } from "echarts";
import { onMounted, nextTick, type Ref, onUnmounted } from "vue";
import { addResizeListener, removeResizeListener } from "howtools";

export default function useECharts(
  el: Ref<HTMLDivElement>,
  theme: "light" | "dark" | "default" = "light"
) {
  //echarts图实例
  let echartInstance: ECharts | null = null;

  //设置默认样式数据
  const defaultOption: echarts.EChartsOption = {
    backgroundColor: theme == "dark" ? "rgba(0,0,0,0)" : "rgba(255,555,255)",
  };

  function addDefaultOption(option: echarts.EChartsOption) {
    Object.assign(defaultOption, option);
  }

  async function setOption(option: echarts.EChartsOption) {
    if (!el.value) {
      await nextTick();
      echartInstance = echarts.init(el.value, theme);
    }
    if (!echartInstance) throw new Error("echarts 实例没有创建成功");

    echartInstance?.setOption(Object.assign(defaultOption, option), true);
  }

  function resize() {
    echartInstance?.resize();
  }

  //初始化echarts图
  onMounted(() => {
    if (!el.value) return;
    echartInstance = echarts.init(el.value, theme);
    addResizeListener(el.value, resize);
  });

  onUnmounted(() => {
    removeResizeListener(el.value, resize);
  });

  return {
    addDefaultOption,
    setOption,
    resize,
    echartInstance,
  };
}
