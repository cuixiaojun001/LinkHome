<template>
  <div class="app-container">
    <!-- 列表 -->
    <el-table :data="list" border stripe>
      <el-table-column label="序号" width="50">
        <template slot-scope="scope">
          {{ (page - 1) * limit + scope.$index + 1 }}
        </template>
      </el-table-column>

      <el-table-column prop="real_name" label="用户姓名" align="center" />
      <el-table-column prop="mobile" label="手机号" align="center" />
      <el-table-column prop="id_card" label="身份证号" align="center" />
      <el-table-column prop="auth_status" label="状态" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" v-if="scope.row.auth_status === 'auditing'">
            审核中
          </el-tag>
          <el-tag type="primary" v-if="scope.row.auth_status === 'authorized'">
            已实名
          </el-tag>
          <el-tag type="danger" v-if="scope.row.auth_status === 'unapprove'">
            审核不通过
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="auth_apply_ts" label="申请时间" align="center" />
      <el-table-column label="操作" align="center" width="200">
        <template slot-scope="scope">
          <el-button
            type="primary"
            size="mini"
            @click="toAuditDetail(scope.row.user_id, 0)"
          >
            查看
          </el-button>
          <el-button
            v-if="scope.row.auth_status === 'auditing'"
            type="warning"
            size="mini"
            @click="toAuditDetail(scope.row.user_id, 1)"
          >
            审批
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页组件 -->
    <el-pagination
      @size-change="changePageSize"
      @current-change="changeCurrentPage"
      :current-page="page"
      :page-sizes="[5, 10, 15, 20]"
      :page-size="limit"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
      style="padding: 30px 0"
    ></el-pagination>
  </div>
</template>

<script>
import utilsApi from "@/api/utils.js";

export default {
  data() {
    return {
      list: [], // 数据列表
      total: 0, // 数据库中的总记录数
      page: 1, // 默认页码
      limit: 5, // 每页记录数
      searchObj: {}, // 查询条件
    };
  },

  created() {
    this.fetchData();
  },

  methods: {
    // 获取列表数据
    fetchData() {
      // 封装分页参数
      var param = {};
      param.offset = (this.page - 1) * this.limit;
      param.limit = this.limit;
      param.query_params = {
        auth_status__in: ["authorized", "auditing", "unapprove"],
      };
      utilsApi.getUserList(param).then((response) => {
        this.list = response.data.data_list;
        this.total = response.data.total;
        for (var i = 0; i < this.list.length; ++i) {
          // 特殊处理:将申请时间的时间戳转为时间字符串
          var time = this.list[i].auth_apply_ts;
          var timeStr = new Date(parseInt(time) * 1000)
            .toLocaleString()

            .replace(/:\d{1,2}$/, " ");
          this.list[i].auth_apply_ts = timeStr;
        }
      });
    },
    // 跳转至审核详情界面
    toAuditDetail(userId, tag) {
      // 若tag为0，则为查看操作
      // 若tag为1，则为审批操作
      this.$router.push({
        path: "/auth/detail",
        query: {
          tag: tag,
          userId: userId,
        },
      });
    },
    // 当修改每页显示的记录数时调用此函数
    changePageSize(size) {
      this.limit = size;
      // 重新获取一次数据
      this.fetchData();
    },
    // 当点击页码进行页面切换时调用此函数
    changeCurrentPage(page) {
      this.page = page;
      // 重新获取一次数据
      this.fetchData();
    },
    // 清空查询条件
    resetData() {
      // 还原查询表单
      this.searchObj = {};
      // 重新获取一次数据
      this.fetchData();
    },
  },
};
</script>
