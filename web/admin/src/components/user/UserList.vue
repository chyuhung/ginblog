<template>
  <div>
    <!-- <h3>Userlist page</h3> -->
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search v-model="queryParam.username" placeholder="输入用户名查找" enter-button allowClear @search="getUserList" />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">新增</a-button>
        </a-col>
      </a-row>
      <a-row>
        <a-table
          :row-key="(data) => data.username"
          :columns="columns"
          :pagination="pagination"
          :data-source="userlist"
          @change="handleTableChange"
          bordered
        >
          <span slot="role" slot-scope="role">{{ role == 1 ? '管理员' : '订阅者' }}</span>
          <template slot="action" slot-scope="data">
            <div class="actionSlot" slot="action">
              <a-button type="primary" style="margin-right: 15px">编辑</a-button>
              <a-button type="danger" @click="deleteUser(data.ID)">删除</a-button>
            </div>
          </template>
        </a-table>
      </a-row>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'role' }
  },
  {
    title: '操作',
    width: '20%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        current: 1,
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        showQuickJumper: true
      },
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagenum: 1,
        pagesize: 5
      }
    }
  },
  created() {
    this.getUserList()
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.userlist = res.data
      this.pagination.total = res.total
    },
    handleTableChange(pagination, filters, sorter) {
      const { current, pageSize } = pagination

      if (pageSize !== this.pagination.pageSize) {
        pagination.current = 1
      }

      this.pagination = pagination
      this.queryParam.pagenum = current
      this.queryParam.pagesize = pageSize

      this.getUserList()
    },
    // 删除用户
    deleteUser(id) {
      this.$confirm({
        title: '警告：确定删除该用户吗？',
        content: '删除后无法恢复',
        onOk: async () => {
          const res = await this.$http.delete(`/user/${id}`)
          if (res.status !== 200) return this.$http.message.error(res.message)
          this.$message.success('已删除')
          this.getUserList()
        },
        onCancel() {}
      })
    }
  }
}
</script>
