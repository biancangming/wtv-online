<template>
  <a-tabs v-model:activeKey="activeKey">
    <a-tab-pane key="1" tab="使用中">
      <a-table size="small" :columns="columns" :dataSource="resultDatas.data">
        <template #emptyText>
          <a-empty>
            暂无数据，<a-button type="link" @click="add">新增</a-button>一个吧
          </a-empty>
        </template>
        <template #bodyCell="{ column, record }">
          <template v-if="column.dataIndex === 'operation'">
            <a-space>
              <a-button size="small" type="primary" @click="copyText(record)">
                复制链接
              </a-button>
              <a-button size="small" type="info" @click="editRouter(record)">
                修改
              </a-button>
              <a-button size="small" type="danger" @click="discardText(record)">
                废弃
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
      <div style="text-align: center; margin-top: 10px">
        <a-button type="primary" block @click="add">新增一条</a-button>
      </div>
    </a-tab-pane>
    <a-tab-pane key="2" tab="废弃" force-render>
      <a-table
        size="small"
        :columns="columns"
        :dataSource="resultRemoveDatas.data"
      >
        <template #emptyText>
          <a-empty> 暂无数据，您还没有废弃的数据 </a-empty>
        </template>
        <template #bodyCell="{ column, record }">
          <template v-if="column.dataIndex === 'operation'">
            <a-space>
              <a-button size="small" type="danger" @click="reOpenText(record)">
                恢复使用
              </a-button>
            </a-space>
          </template>
        </template>
      </a-table>
    </a-tab-pane>
  </a-tabs>
</template>
<script lang="ts" setup>
import { useShareUrls, useShareUpdateStatus } from "@/api/share";
import { copyToClipboard } from "howtools";
import { useRouter } from "vue-router";
import { message } from "ant-design-vue";
import { ref } from "vue";

const activeKey = ref("1");

const { push } = useRouter();
const columns = [
  {
    title: "链接ID",
    dataIndex: "uuid",
  },
  // {
  //   title: "更新时间",
  //   dataIndex: "updateTime",
  // },
  {
    title: "文件类型",
    dataIndex: "fileType",
  },
  {
    title: "操作",
    dataIndex: "operation",
  },
];
const { data: resultDatas, reload: executeDatas } = useShareUrls(true, 0);
const { data: resultRemoveDatas, reload: executeRemoveDatas } = useShareUrls(
  true,
  1
);
const { execute: executeUpdateStatus } = useShareUpdateStatus();

function editRouter(record: { uuid: string }) {
  push({
    name: "edit",
    query: {
      uuid: record.uuid,
    },
  });
}

function copyText(record: { uuid: string; fileType: string }) {
  if (!record.fileType) {
    message.error("复制错误，文件类型错误，可联系管理员");
    return;
  }
  const url = `${location.origin}/api/share/${record.uuid}.${record.fileType}`;
  copyToClipboard(url).then(() => {
    message.success("复制到剪贴板成功");
  });
}

function add() {
  push({
    name: "edit",
  });
}

// 废弃文件
function discardText(record: Record<string, string>) {
  executeUpdateStatus({
    data: {
      ...record,
      useStatus: 1,
    },
  }).then((res) => {
    if (res.code == 0) {
      message.success("废弃成功");
      executeDatas();
      executeRemoveDatas();
    } else {
      message.error("联系管理员");
    }
  });
}

// 启用文件
function reOpenText(record: Record<string, string>) {
  executeUpdateStatus({
    data: {
      ...record,
      useStatus: 0,
    },
  }).then((res) => {
    if (res.code == 0) {
      message.success("还原使用成功");
      executeDatas();
      executeRemoveDatas();
    } else {
      message.error("联系管理员");
    }
  });
}
</script>
