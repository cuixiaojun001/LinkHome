<template>
  <div class="app-container">
    <!-- 查询表单 -->
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item label="手机号">
        <el-input v-model="searchObj.phone" placeholder="手机号" />
      </el-form-item>

      <el-button type="primary" icon="el-icon-search" @click="fetchData()">
        查询
      </el-button>
      <el-button type="default" @click="resetData()">清空</el-button>
    </el-form>
    <!-- 列表 -->
    <el-table :data="list" border stripe>
      <el-table-column label="#" width="50" fixed="left">
        <template slot-scope="scope">
          {{ (page - 1) * limit + scope.$index + 1 }}
        </template>
      </el-table-column>

      <el-table-column prop="userName" label="用户姓名" align="center" />
      <el-table-column prop="city" label="所在城市" align="center" />
      <el-table-column
        prop="houseType"
        label="房屋类型"
        align="center"
        width="150px"
      />
      <el-table-column prop="TrafficInfoJson" label="交通要求" align="center" />
      <el-table-column
        prop="houseFacilityJson"
        label="房屋设施要求"
        align="center"
        width="150px"
      />
      <el-table-column
        prop="moneyBudget"
        label="金额预算"
        align="center"
        width="150px"
      />
      <el-table-column prop="floor" label="楼层" align="center" />
      <el-table-column prop="commutingTime" label="通勤时间" align="center" />
      <el-table-column
        prop="commanyAddress"
        label="公司地址"
        align="center"
        width="150px"
      />
      <el-table-column
        prop="desiredResidenceArea"
        label="期望居住地区"
        align="center"
        width="150px"
      />
      <el-table-column label="操作" align="center" width="200" fixed="right">
        <template style="margin-top: 10px"
          ><el-button type="warning" size="mini" @click="match">
            匹配房源
          </el-button>
          <br />
          <el-badge
            style="margin-top: 5px"
            :value="2"
            class="item"
            type="warning"
          >
            <el-button size="small">房源信息</el-button>
          </el-badge>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      title="正在匹配房源..."
      :visible.sync="dialogMatchVisible"
      width="30%"
    >
      <el-progress
        v-if="pt !== 100"
        type="circle"
        :percentage="pt"
        style="margin-left: 135px"
      ></el-progress>
      <el-progress
        v-else
        type="circle"
        :percentage="100"
        status="success"
        style="margin-left: 135px"
      ></el-progress>
    </el-dialog>

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
export default {
  data() {
    return {
      list: null, // 数据列表
      total: 0, // 数据库中的总记录数
      page: 1, // 默认页码
      limit: 5, // 每页记录数
      searchObj: {}, // 查询条件
      pt: 0, // 环形进度条的百分比
      dialogMatchVisible: false, // 控制环形进度条是否显示
    };
  },

  created() {
    this.fetchData();
  },

  methods: {
    // 获取列表数据
    fetchData() {
      this.list = [
        {
          id: 1,
          userName: "刘土鳖",
          city: "深圳",
          houseType: "住宅/公寓",
          TrafficInfoJson: "略",
          houseFacilityJson: "略",
          moneyBudget: "10000以内",
          floor: "6-10层",
          commutingTime: "10分钟",
          commanyAddress: "深圳xxx街道xxx号",
          desiredResidenceArea: "深圳",
        },
      ];
    },
    // 匹配房源
    match() {
      this.dialogMatchVisible = true;
      // 开始滚动进度条
      this.pt = 1;
      let id = setInterval(() => {
        this.pt += 1;
        if (this.pt === 100) {
          clearInterval(id);
          // 延时1秒后关闭进度条
          setTimeout(() => {
            this.dialogMatchVisible = false;
            this.$notify({
              title: "成功",
              message: "匹配完成",
              type: "success",
            });
          }, 1000);
        }
      }, 100);

      // 首先匹配房屋类型
      // 匹配期望居住地址
      // 匹配楼层
      // 匹配预算
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
