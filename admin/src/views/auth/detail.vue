<template>
  <div class="app-container">
    <el-form label-width="100px" class="form-table">
      <el-row>
        <el-col :span="6">
          <el-form-item label="状态">
            <span v-if="user.auth_status === 'auditing'">审核中</span>
            <span v-if="user.auth_status === 'authorized'">已实名</span>
            <span v-if="user.auth_status === 'unapprove'">审核不通过</span>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="申请时间">
            {{ user.auth_apply_ts }}
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <el-form-item label="姓名">
            {{ user.real_name }}
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="性别">
            <span v-if="user.gender === 'male'">男</span>
            <span v-if="user.gender === 'female'">女</span>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="年龄">
            {{ user.age }}
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="手机">
            {{ user.mobile }}
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="6">
          <el-form-item label="身份证号">
            {{ user.id_card }}
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="24">
          <el-form-item label="身份证正面">
            <span>
              <el-image style="width: 150px" :src="user.id_card_front" />
            </span>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="24">
          <el-form-item label="身份证反面">
            <span>
              <el-image style="width: 150px" :src="user.id_card_back" />
            </span>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row style="text-align: center" v-if="tag === 0">
        <el-button @click="back"> 返回 </el-button>
      </el-row>
    </el-form>

    <el-form label-width="170px">
      <el-form-item label="是否通过" v-if="tag === 1">
        <el-radio-group v-model="auth_status_flag">
          <el-radio label="1"> 通过 </el-radio>
          <el-radio label="0"> 不通过 </el-radio>
        </el-radio-group>
      </el-form-item>

      <el-row style="text-align: center" v-if="tag === 1">
        <el-button type="primary" @click="onSubmit()"> 确定 </el-button>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import utilsApi from "@/api/utils.js";

export default {
  data() {
    return {
      tag: 0, // 操作标识
      saveBtnDisabled: false, //防止重复提交
      user: {},
      userId: 0,
      auth_status_flag: "1",
    };
  },

  mounted() {
    // 获取实名认证列表传递过来的参数
    this.tag = this.$route.query.tag;
    this.userId = this.$route.query.userId;
    this.fetchDataById();
  },

  methods: {
    // 获取指定用户id的审核详情信息
    // 根据id查询用户详情
    fetchDataById() {
      utilsApi.getUserDetailById(this.userId).then((response) => {
        this.user = response.data;
        // 特殊处理:将注册时间的时间戳转为时间字符串
        var time = this.user.auth_apply_ts;
        var timeStr = new Date(parseInt(time) * 1000)
          .toLocaleString()
          .replace(/:\d{1,2}$/, " ");
        this.user.auth_apply_ts = timeStr;
      });
    },
    // 返回实名认证列表
    back() {
      this.$router.push("/auth/index");
    },
    // 提交审核
    onSubmit() {
      // 判断审核是否通过
      if (this.auth_status_flag === "0") {
        this.user.auth_status = "unapprove";
      } else if (this.auth_status_flag === "1") {
        this.user.auth_status = "authorized";
      }
      // 更新用户信息
      utilsApi.updateUserDetailById(this.userId, this.user).then((response) => {
        if (response.message === "SUCCESS") {
          this.$message({
            message: "审核完成",
            type: "success",
          });
        }
      });
      // 跳转回实名认证列表
      this.$router.push("/auth/index");
    },
  },
};
</script>
