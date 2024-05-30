<template>
  <div style="margin: 20px">
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="收租日期">
        <span>{{ form.szrq }}</span> </el-form-item
      ><el-form-item label="账单周期">
        <el-date-picker
          v-model="form.zdzq"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        >
        </el-date-picker>
      </el-form-item>
      <el-form-item label="基本费用">
        <span>
          <b>租金:{{ form.zj }}</b>
        </span>
        <span style="margin-left: 20px">
          <b>押金:{{ form.yj }}</b>
        </span>
      </el-form-item>
      <el-form-item label="其它费用"> </el-form-item>
      <el-table :data="tableData" style="width: 100%" size="small">
        <el-table-column prop="cbfyx" label="抄表费用项" width="100">
          <template slot-scope="scope">
            <span>{{ scope.row.cbfyx }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="bds" label="表读数" width="100" />
        <el-table-column prop="yl" label="用量" width="100" />
        <el-table-column prop="dj" label="单价" width="100" />
        <el-table-column prop="je" label="金额" width="100" />
        <el-table-column prop="cbrq" label="抄表日期" width="100" />
      </el-table>
      <template v-for="(data, index) in otherData">
        <span :key="index">{{ data.name }}</span>
      </template>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        szrq: "2022-04-22",
        zdzq: "",
        zj: "10000.00",
        yj: "10000.00",
      },
      tableData: [
        {
          cbfyx: "电费",
          bds: "100 -- 100",
          yl: "0.00",
          dj: "1.00",
          je: "0.00",
        },
        {
          cbfyx: "水费",
          bds: "100 -- 100",
          yl: "0.00",
          dj: "1.00",
          je: "0.00",
        },
      ],
      otherData: [
        {
          name: "垃圾处理费",
          feed: "50",
        },
        {
          name: "网费",
          feed: "100",
        },
        {
          name: "物业管理费",
          feed: "50",
        },
      ],
    };
  },
};
</script>

<style></style>
