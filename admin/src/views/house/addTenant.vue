<template>
  <div>
    <el-dialog title="添加租客" :visible.sync="dialogFormVisible" fullscreen>
      <hr />
      <el-form :model="form" inline label-width="90px">
        <table>
          <tr>
            <td>
              <el-form-item
                label="租客姓名"
                size="medium"
                style="margin-top: -10px"
              >
                <el-input></el-input>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td>
              <el-form-item
                label="电话号码"
                size="medium"
                style="margin-top: -30px"
              >
                <el-input></el-input>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td>
              <el-form-item
                label="身份证号"
                size="medium"
                style="margin-top: -30px"
              >
                <el-input></el-input>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td>
              <el-form-item
                label="租期"
                size="medium"
                style="margin-top: -30px"
              >
                <el-date-picker
                  v-model="value1"
                  type="daterange"
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                >
                </el-date-picker
              ></el-form-item>
            </td>
          </tr>
          <tr>
            <td>
              <el-form-item
                label="电费"
                size="medium"
                style="margin-top: -20px"
              >
                <template>
                  <el-input
                    style="width: 100px"
                    placeholder="请输入价格"
                  ></el-input>
                  <el-input
                    placeholder="元/度"
                    disabled
                    style="width: 70px"
                  ></el-input>
                  <el-input
                    style="width: 100px; margin-left: 30px"
                    placeholder="初始读数"
                  ></el-input>
                </template>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td>
              <el-form-item
                label="水费"
                size="medium"
                style="margin-top: -20px"
              >
                <template>
                  <el-input
                    style="width: 100px"
                    placeholder="请输入价格"
                  ></el-input>
                  <el-input
                    placeholder="元/吨"
                    disabled
                    style="width: 70px"
                  ></el-input>
                  <el-input
                    style="width: 100px; margin-left: 30px"
                    placeholder="初始读数"
                  ></el-input>
                </template>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td>
              <el-form-item
                label="物业管理费"
                size="medium"
                style="margin-top: -20px"
              >
                <template>
                  <el-input
                    style="width: 100px"
                    placeholder="请输入价格"
                  ></el-input>
                  <el-input
                    placeholder="元"
                    disabled
                    style="width: 50px"
                  ></el-input>
                </template>
              </el-form-item>
            </td>
          </tr>
        </table>
        <el-form-item style="float: right">
          <el-button type="primary" @click="addTenant">添加</el-button>
        </el-form-item></el-form
      ></el-dialog
    >
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialogFormVisible: true,
      form: {},
    };
  },
  methods: {
    addTenant() {
      // 添加租客后生成订单

      // 跳转至订单界面
      this.$router.push("/order/order");
    },
  },
};
</script>

<style></style>
