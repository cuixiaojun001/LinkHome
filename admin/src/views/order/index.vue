<template>
  <div class="app-container">
    <!-- 查询表单 -->
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item label="选择房源">
        <el-select v-model="selectHouse" placeholder="请选择" clearable>
          <el-option
            v-for="(order, index) in orderList"
            :key="index"
            :label="order.house_info.title"
            :value="order.house_id"
          />
        </el-select>
      </el-form-item>
    </el-form>
    <!-- 列表 -->
    <el-table :data="orderDetail" border stripe>
      <el-table-column label="#" width="50" fixed="">
        <template slot-scope="scope">
          {{ (page - 1) * limit + scope.$index + 1 }}
        </template>
      </el-table-column>

      <el-table-column prop="zh" label="租户" width="200px" align="center">
        <template slot-scope="scope">
          <span style="color: blue">{{ scope.row.house_info.title }}</span>
          -
          <span style="color: blue">{{ scope.row.user_info.real_name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        prop="zdlx"
        label="账单类型"
        width="100px"
        align="center"
      >
        <template>
          <span style="color: orange">租金账单</span>
        </template>
      </el-table-column>
      <el-table-column
        prop="start_date"
        label="收租日"
        width="100px"
        align="center"
      />
      <el-table-column
        prop="zdzq"
        label="账单周期"
        width="200px"
        align="center"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.start_date }}</span>
          至
          <span>{{ scope.row.end_date }}</span>
        </template></el-table-column
      >
      <el-table-column prop="xj" label="小计" align="center">
        <template slot-scope="scope">
          <span style="color: red">¥{{ scope.row.pay_money }}</span>
        </template></el-table-column
      >
      <el-table-column prop="fym" label="费用名" align="left" width="100px">
        <template slot-scope="scope">
          <span
            v-for="(f, index) in scope.row.fym"
            :key="index"
            style="display: block"
          >
            {{ f }}
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="je" label="金额" align="left">
        <template slot-scope="scope">
          <span
            v-for="(j, index) in scope.row.je"
            :key="index"
            style="display: block; color: red"
          >
            ¥{{ j }}
          </span>
        </template></el-table-column
      >
      <el-table-column prop="bz" label="备注" align="center" />
      <el-table-column
        prop="cb"
        label="抄表"
        align="center"
        @click="meterReading"
      />
      <el-table-column label="操作" align="center" width="150" fixed="right">
        <template>
          <el-link type="primary">账单详情</el-link>
          <el-link type="primary" style="margin-left: 20px">到账</el-link>
          <br />
          <el-link type="primary" @click="meterReading">抄表</el-link>
          <el-link type="primary" style="margin-left: 20px">更多</el-link>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :title="meterTitle" :visible.sync="dialogMeterVisible">
      <el-table :data="meterData">
        <el-table-column
          property="fyx"
          label="费用项"
          align="center"
        ></el-table-column>
        <el-table-column
          property="sqds"
          label="上期读数"
          align="center"
        ></el-table-column>
        <el-table-column property="name" label="本期读数" align="center">
          <template> <el-input></el-input> </template
        ></el-table-column>
        <el-table-column
          property="fy"
          label="费用"
          align="center"
        ></el-table-column>
      </el-table>
      <el-form>
        <el-form-item
          label="抄表日期"
          style="margin-top: 30px; margin-left: 30px"
        >
          <el-date-picker type="date" placeholder="选择日期" align="center">
          </el-date-picker
        ></el-form-item>
        <el-form-item style="text-align: center">
          <el-button type="primary" size="medium">保存</el-button>
        </el-form-item>
      </el-form>
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
import utilsApi from "@/api/utils.js";

export default {
  data() {
    return {
      orderDetail: [], // 订单数据
      orderList: [], // 数据列表
      total: 0, // 数据库中的总记录数
      page: 1, // 默认页码
      limit: 5, // 每页记录数
      searchObj: {}, // 查询条件
      dialogMeterVisible: false, // 控制抄表对话框是否显示
      meterTitle: "", // 抄表对话框标题
      meterData: [], // 抄表显示数据
      selectHouse: "", // 选择的房源
    };
  },

  created() {
    this.fetchData();
  },

  watch: {
    // 监听选择了哪个房源
    selectHouse(newVal, oldVal) {
      // newVal为当前选中的房源，显示该房源订单数据
      var order = this.orderList[newVal - 1];
      // 每次选择前清空一下数据
      this.orderDetail = [];
      this.orderDetail.push(order);
    },
  },

  methods: {
    // 获取列表数据
    fetchData() {
      // 封装分页参数
      var param = {};
      param.offset = (this.page - 1) * this.limit;
      param.limit = this.limit;
      param.query_params = {};
      utilsApi.getOrderList(param).then((response) => {
        this.orderList = response.data.data_list;
        this.total = response.data.total;
        // 封装费用名数组
        for (var i = 0; i < this.orderList.length; ++i) {
          var je = [];
          var fym = [];
          je.push(this.orderList[i].house_info.rent_money);
          je.push(
            this.orderList[i].house_info.deposit_ratio *
              this.orderList[i].house_info.rent_money
          );
          je.push(this.orderList[i].house_info.water_rent);
          je.push(this.orderList[i].house_info.electricity_rent);
          je.push(this.orderList[i].house_info.strata_fee);
          fym.push("租金");
          fym.push("押金");
          fym.push("水费");
          fym.push("电费");
          fym.push("管理费");
          this.orderList[i].je = je;
          this.orderList[i].fym = fym;
        }
        // 显示第一个房源数据
        this.selectHouse = this.orderList[0].house_id;
      });
    },
    // fetchData() {
    //   this.orderList = [
    //     {
    //       id: 1,
    //       zh: "房源1 - 张三",
    //       zdlx: "租金账单",
    //       szr: "2022-04-21",
    //       zdzq: "2022-04-20 至 2022-05-19",
    //       xj: "15203.00",
    //       fym: ["押金", "租金", "电费", "水费", "垃圾处理费", "网费", "物业费"],
    //       je: [
    //         "10000.00",
    //         "5000.00",
    //         "100.00",
    //         "100.00",
    //         "1.00",
    //         "1.00",
    //         "1.00",
    //       ],
    //       bz: "无",
    //       cb: "",
    //     },
    //   ];
    // },

    // 显示抄表对话框
    meterReading() {
      // 拼接抄表对话框标题
      this.meterTitle = "抄表";
      this.dialogMeterVisible = true;
      // 准备抄表显示数据
      this.meterData = [
        { fyx: "电费", sqds: "0", fy: "0.00" },
        { fyx: "水费", sqds: "0", fy: "0.00" },
      ];
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
