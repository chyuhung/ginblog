<template>
  <div>
    <h3>Userlist page</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="输入用户名查找" enter-button />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">新增</a-button>
        </a-col>
      </a-row>
      <a-row>
        <a-table></a-table>
      </a-row>
    </a-card>
  </div>
</template>
