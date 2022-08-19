<template>
  <div class="wrapper">
    <a-row>
      <a-col :offset="2">
        <a-button type="primary" @click="backHome"
          ><left-outlined />返回列表页面</a-button
        >
      </a-col>
    </a-row>
    <a-spin tip="Loading..." :spinning="contentEditSpinning">
      <a-form :label-col="{ span: 2 }" style="margin-top: 10px">
        <a-form-item label="文本内容" v-bind="validateInfos.content">
          <a-textarea
            v-model:value="modelRef.content"
            :auto-size="{ minRows: 20, maxRows: 25 }"
          ></a-textarea>
        </a-form-item>
        <a-form-item label="文件类型" v-bind="validateInfos.fileType">
          <a-radio-group v-model:value="modelRef.fileType">
            <a-radio value="m3u">m3u</a-radio>
            <a-radio value="txt">txt</a-radio>
            <a-radio value="json">json</a-radio>
            <a-radio value="yml">yml</a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 14, offset: 2 }">
          <a-space>
            <a-button type="primary" @click.prevent="onSubmit">保存</a-button>
            <a-button @click="resetFields"> 清空填写内容 </a-button>
            <a-button @click="backHome"><left-outlined />返回列表页面</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-spin>
    <a-row
      style="margin-bottom: 10px"
      v-if="modelRef.uuid && modelRef.fileType"
    >
      <a-col :offset="2">
        <div>你可以通过以下链接查看生成的内容,访问之前务必点击保存生成链接</div>
        <div>
          {{
            `${location.origin}/api/share/${modelRef.uuid}.${modelRef.fileType}`
          }}
        </div>
      </a-col>
    </a-row>
  </div>
</template>
<script lang="ts" setup>
import { reactive, toRaw, onMounted, ref } from "vue";
import { Form } from "ant-design-vue";
import { useShareGet, useShareUpdate } from "@/api/share";
import { useDataTip } from "@/utils/requests";
import { useRoute, useRouter } from "vue-router";
import { LeftOutlined } from "@ant-design/icons-vue";
const useForm = Form.useForm;
const location = window.location;
const contentEditSpinning = ref(false);

const route = useRoute();
const { execute: executeShareUpdate, data, error } = useShareUpdate(); // 新增或者分享更新的内容
const { execute: executeGet } = useShareGet(); // 查看
useDataTip(data, error);

onMounted(() => {
  if (!route.query.uuid) return;
  contentEditSpinning.value = true;
  executeGet({
    params: {
      uuid: route.query.uuid,
    },
  }).then((res) => {
    const { uuid, content, fileType } = (res.data || {}) as any;
    modelRef.uuid = uuid;
    modelRef.content = content;
    modelRef.fileType = fileType;
    contentEditSpinning.value = false;
  });
});

const modelRef = reactive({
  uuid: "",
  public: false,
  content: "",
  fileType: "",
});
const rulesRef = reactive({
  fileType: [
    {
      required: true,
      message: "请选择一个文件类型",
    },
  ],
  content: [
    {
      required: true,
      message: "请填入内容",
    },
    {
      min: 50,
      message: "内容不少于50字",
    },
    {
      max: 145000,
      message: "最多输入14.5万个字符, 有疑问公众号私信站长哦",
    },
  ],
});
const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef, {
  onValidate: (...args) => console.log(...args),
});
const onSubmit = () => {
  validate()
    .then(() => {
      console.log(toRaw(modelRef));
      executeShareUpdate({
        data: {
          ...modelRef,
          type: undefined,
        },
      }).then((res: any) => {
        const { uuid } = res.data || {};
        if (!uuid) return;
        modelRef.uuid = uuid;
      });
    })
    .catch((err) => {
      console.log("error", err);
    });
};

const router = useRouter();
function backHome() {
  router.push({ name: "home" });
}
</script>
<style lang="less" scoped>
.wrapper {
  max-width: 1200px;
  margin: 0 auto;
}
</style>
