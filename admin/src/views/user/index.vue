<template>
  <div class="app-container">
    <!-- 查询表单 -->
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item label="手机号">
        <el-input v-model="searchObj.mobile" placeholder="手机号" />
      </el-form-item>

      <el-form-item label="用户状态">
        <el-select v-model="searchObj.state" placeholder="请选择" clearable>
          <el-option label="正常" value="normal" />
          <el-option label="锁定" value="deleted" />
        </el-select>
      </el-form-item>
      <el-button type="primary" icon="el-icon-search" @click="fetchData()">
        查询
      </el-button>
      <el-button type="default" @click="resetData()">清空</el-button>
    </el-form>
    <!-- 列表 -->
    <el-table :data="list" border stripe>
      <el-table-column label="#" width="60">
        <template slot-scope="scope">
          {{ (page - 1) * limit + scope.$index + 1 }}
        </template>
      </el-table-column>

      <el-table-column
        prop="username"
        label="用户姓名"
        align="center"
        sortable
      />
      <el-table-column prop="mobile" label="手机号" align="center" sortable />
      <el-table-column prop="gender" label="性别" align="center">
        <template slot-scope="scope">
          <el-tag type="info" v-if="scope.row.gender === 'male'">男</el-tag>
          <el-tag type="info" v-if="scope.row.gender === 'female'">女</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="role" label="用户角色" align="center">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.role === 'admin'">管理员</el-tag>
          <el-tag v-if="scope.row.role === 'tenant'">租客</el-tag>
          <el-tag v-if="scope.row.role === 'landlord'">房东</el-tag>
          <el-tag v-if="scope.row.role === 'steward'">管家</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="用户状态" width="90" align="center">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.state === 'normal'" type="success">
            正常
          </el-tag>
          <el-tag v-else type="danger"> 锁定 </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        prop="create_ts"
        label="注册时间"
        align="center"
        sortable
      />
      <el-table-column label="操作" align="center" width="200">
        <template slot-scope="scope">
          <el-button
            v-if="scope.row.state == 'normal'"
            type="warning"
            size="mini"
            @click="lock(scope.row.id, 0)"
          >
            锁定
          </el-button>
          <el-button
            v-else
            type="danger"
            size="mini"
            @click="lock(scope.row.id, 1)"
          >
            解锁
          </el-button>
          <el-button
            type="primary"
            size="mini"
            @click="toDetail(scope.row.user_id)"
          >
            查看详情
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
      list: null, // 数据列表
      total: 0, // 数据库中的总记录数
      page: 1, // 默认页码
      limit: 5, // 每页记录数
      searchObj: {
        mobile: null,
        state: null,
      }, // 查询条件
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
      param.query_params = {};
      if (this.searchObj.mobile !== null && this.searchObj.mobile !== "") {
        param.query_params.mobile = this.searchObj.mobile;
      }
      if (this.searchObj.state !== null && this.searchObj.state !== "") {
        param.query_params.state = [];
        param.query_params.state.push(this.searchObj.state);
      }
      utilsApi.getUserList(param).then((response) => {
        this.list = response.data.data_list;
        this.total = response.data.total;
        for (var i = 0; i < this.list.length; ++i) {
          // 特殊处理:将注册时间的时间戳转为时间字符串
          var time = this.list[i].create_ts;
          var timeStr = new Date(parseInt(time) * 1000)
            .toLocaleString()

            .replace(/:\d{1,2}$/, " ");
          this.list[i].create_ts = timeStr;
        }
      });
    },

    // 跳转至用户信息详情页
    toDetail(id) {
      this.$router.push({
        path: "/user/detail",
        query: {
          userId: id,
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
