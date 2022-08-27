<template>
  <a-card title="登录">
    <a-form
      :model="formState"
      name="basic"
      :label-col="{ span: 3 }"
      :wrapper-col="{ span: 10 }"
      autocomplete="off"
      @finish="onFinish"
      @finishFailed="onFinishFailed"
    >
      <a-form-item
        label="用户名"
        name="username"
        :rules="[{ required: true, message: '请输入用户名' }]"
      >
        <a-input v-model:value="formState.username" />
      </a-form-item>

      <a-form-item
        label="密码"
        name="password"
        :rules="[{ required: true, message: '请输入密码' }]"
      >
        <a-input-password v-model:value="formState.password" />
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 3, span: 16 }">
        <a-button type="primary" html-type="submit">登录</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>
<script lang="ts">
import { UseLogin } from "@/api/config";
import { defineComponent, reactive } from "vue";
import { message } from "ant-design-vue";

interface FormState {
  username: string;
  password: string;
}
export default defineComponent({
  setup() {
    const { execute: reloadLogin } = UseLogin();
    const formState = reactive<FormState>({
      username: "",
      password: "",
    });

    const onFinish = (values: any) => {
      reloadLogin({
        data: values,
      }).then((res) => {
        const logined = res == "1";
        sessionStorage.setItem("logined", JSON.stringify(logined));
        if (logined) {
          location.reload();
        } else {
          message.error("账户名或者密码错误");
        }
      });
      console.log("Success:", values);
    };

    const onFinishFailed = (errorInfo: any) => {
      console.log("Failed:", errorInfo);
    };
    return {
      formState,
      onFinish,
      onFinishFailed,
    };
  },
});
</script>
