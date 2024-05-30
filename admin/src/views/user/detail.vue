<template>
  <div class="app-container">
    <el-form label-width="100px" class="form-table">
      <el-row>
        <el-col :span="6">
          <el-form-item label="用户名">
            <span v-if="isEdit === false">
              {{ userDetail.username }}
            </span>
            <el-input v-else v-model="userDetail.username"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="用户角色">
            <span>
              <span v-if="userDetail.role === 'admin'">管理员</span>
              <span v-if="userDetail.role === 'tenant'">租客</span>
              <span v-if="userDetail.role === 'landlord'">房东</span>
              <span v-if="userDetail.role === 'steward'">管家</span>
            </span>
            <!-- <el-select v-else v-model="userDetail.role" placeholder="用户角色">
              <el-option value="admin" label="管理员"></el-option>
              <el-option value="tenant" label="租客"></el-option>
              <el-option value="landlord" label="房东"></el-option>
              <el-option value="steward" label="管家"></el-option>
            </el-select> -->
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <el-form-item label="手机号">
            {{ userDetail.mobile }}
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="真实姓名">
            {{ userDetail.real_name }}
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="邮箱">
            <span v-if="isEdit === false">{{ userDetail.mail }}</span>
            <el-input v-else v-model="userDetail.mail"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="头像">
            <el-image
              style="width: 150px"
              :src="userDetail.avatar"
              :preview-src-list="[userDetail.avatar]"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <el-form-item label="用户状态">
            <span v-if="userDetail.state === 'normal'">正常</span>
            <span v-if="userDetail.state === 'deleted'">已锁定</span>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="身份证号">
            {{ userDetail.id_card }}
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="职业">
            <span v-if="isEdit === false">{{ userDetail.career }}</span>
            <el-input v-else v-model="userDetail.career"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="性别">
            <span v-if="isEdit === false">
              <span v-if="userDetail.gender === 'male'">男</span>
              <span v-if="userDetail.gender === 'female'">女</span>
            </span>
            <el-radio-group v-else v-model="userDetail.gender">
              <el-radio label="male">男</el-radio>
              <el-radio label="female">女</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <el-form-item label="注册时间">
            {{ userDetail.create_ts }}
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="用户简介">
            <span v-if="isEdit === false">{{ userDetail.userDesc }}</span>
            <el-input v-else v-model="userDetail.userDesc"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="爱好">
            <span v-if="isEdit === false">{{ userDetail.hobby }}</span>
            <el-input v-else v-model="userDetail.hobby"></el-input>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>

    <el-form>
      <el-row style="text-align: center">
        <el-button
          v-if="isEdit === false"
          type="warning"
          @click="editUserDetail"
        >
          编辑
        </el-button>
        <el-button v-else type="primary" @click="saveUserDetail"
          >保存</el-button
        >
        <el-button type="info" @click="back"> 返回 </el-button>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import utilsApi from "@/api/utils.js";

export default {
  data() {
    return {
      userDetail: {},
      isEdit: false, // 是否编辑
      userId: 0, // 用户id
    };
  },

  created() {
    // 接收到用户列表传递过来的用户id
    this.userId = this.$route.query.userId;
    this.fetchDataById();
  },

  methods: {
    // 根据id查询用户详情
    fetchDataById() {
      utilsApi.getUserDetailById(this.userId).then((response) => {
        this.userDetail = response.data;
        // 特殊处理:将注册时间的时间戳转为时间字符串
        var time = this.userDetail.create_ts;
        var timeStr = new Date(parseInt(time) * 1000)
          .toLocaleString()

          .replace(/:\d{1,2}$/, " ");
        this.userDetail.create_ts = timeStr;
      });
    },
    // 编辑用户详情
    editUserDetail() {
      this.isEdit = true;
    },

    // 更新用户详情
    saveUserDetail() {
      this.isEdit = false;
      utilsApi
        .updateUserDetailById(this.userId, this.userDetail)
        .then((response) => {
          if (response.message === "SUCCESS") {
            this.$message({
              message: "更新成功",
              type: "success",
            });
          }
        });
    },

    // 返回用户信息列表
    back() {
      this.$router.push("/user/index");
    },
  },
};
</script>
