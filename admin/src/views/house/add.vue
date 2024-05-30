<template>
  <div style="margin: 30px">
    <div v-if="step === 1">
      <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="房源名称">
          <el-input
            v-model="form.title"
            style="width: 400px"
            placeholder="请输入房源名称"
          ></el-input>
        </el-form-item>
        <el-form-item label="房源展示">
          <el-upload
            class="avatar-uploader"
            action="http://3w5q328382.51vip.biz/api/v1/upload/"
            :show-file-list="false"
            :on-success="handleHouseImgSuccess"
          >
            <img v-if="imageUrl" :src="imageUrl" class="avatar" />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="房屋位置">
          <template>
            <el-input
              v-model="form.address"
              style="width: 345px"
              disabled
            ></el-input>
            <el-button icon="el-icon-caret-right" @click="showMap"></el-button>
          </template>
        </el-form-item>
        <el-form-item label="详细地址">
          <template>
            <el-input
              v-model="province"
              placeholder="省"
              style="width: 100px"
            ></el-input>
            <el-input
              v-model="city"
              placeholder="市"
              style="width: 100px; margin-left: 5px"
            ></el-input>
            <el-input
              v-model="district"
              placeholder="县/区"
              style="width: 100px; margin-left: 5px"
            ></el-input>
            <el-input
              v-model="street"
              placeholder="街道"
              style="width: 100px; margin-left: 5px"
            ></el-input>
            <el-input
              v-model="streetNumber"
              placeholder="门牌号"
              style="width: 100px; margin-left: 5px"
            ></el-input>
          </template>
        </el-form-item>
        <el-form-item label="房源类型">
          <el-select v-model="form.house_type">
            <el-option label="住宅" value="department"></el-option>
            <el-option label="小区" value="community"></el-option>
            <el-option label="普通住宅" value="residential"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="收款账号">
          <el-button size="mini" @click="selectPayAccount"
            >选择收款账号</el-button
          >
        </el-form-item>
        <el-form-item>
          <el-button @click="cancel" size="medium">取消</el-button>
          <el-button @click="next" type="primary" size="medium"
            >下一步</el-button
          >
        </el-form-item>
      </el-form>
    </div>

    <div v-if="step === 2">
      <el-form ref="form" :model="form" label-width="90px">
        <el-form-item label="楼栋号">
          <el-input
            v-model="form.loudong"
            style="width: 250px"
            placeholder="请输入楼栋号(没有可以不填)"
          ></el-input>
          栋
        </el-form-item>
        <el-form-item label="房号">
          <el-input
            v-model="form.room_num"
            style="width: 250px"
            placeholder="请输入房号(没有可以不填)"
          ></el-input>
        </el-form-item>
        <el-form-item label="租赁类型">
          <el-select
            v-model="form.rent_type"
            placeholder="请选择"
            style="width: 120px"
          >
            <el-option label="整租" value="whole"></el-option>
            <el-option label="合租" value="share"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="楼层">
          <template>
            <el-select
              v-model="form.floor"
              placeholder="请选择"
              style="width: 120px"
            >
              <el-option
                :key="count"
                v-for="(count, index) in floor"
                :label="count"
                :value="index + 1"
              ></el-option>
            </el-select>
            <el-select
              v-model="form.max_floor"
              placeholder="请选择"
              style="width: 120px; margin-left: 10px"
            >
              <el-option
                :key="count"
                v-for="(count, index) in floorTemp"
                :label="count"
                :value="index + 1"
              ></el-option>
            </el-select>
            <el-radio-group
              v-model="form.has_elevator"
              style="margin-left: 5px"
            >
              <el-radio :label="0">电梯</el-radio>
              <el-radio :label="1">楼梯</el-radio>
            </el-radio-group>
          </template>
        </el-form-item>
        <el-form-item label="户型">
          <el-select
            v-model="form.bedroom_num"
            placeholder="请选择"
            style="width: 80px"
          >
            <el-option
              :key="count"
              v-for="(count, index) in room"
              :label="count"
              :value="index + 1"
            ></el-option>
          </el-select>
          <el-select
            v-model="form.living_room_num"
            placeholder="请选择"
            style="width: 80px; margin-left: 5px"
          >
            <el-option
              :key="count"
              v-for="(count, index) in livingRoom"
              :label="count"
              :value="index + 1"
            ></el-option>
          </el-select>
          <el-select
            v-model="form.toilet_num"
            placeholder="请选择"
            style="width: 80px; margin-left: 5px"
          >
            <el-option
              :key="count"
              v-for="(count, index) in restRoom"
              :label="count"
              :value="index + 1"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="面积">
          <el-input-number
            v-model="form.area"
            controls-position="right"
            @change="handleChange"
            :min="1"
            style="width: 250px"
          ></el-input-number>

          平米</el-form-item
        >
        <el-form-item label="月租金">
          <el-input-number
            v-model="form.rent_money"
            controls-position="right"
            @change="handleChange"
            :min="1"
            style="width: 250px"
          ></el-input-number>

          元</el-form-item
        >
        <el-form-item label="收租周期">
          <el-select
            v-model="form.pay_ratio"
            placeholder="请选择"
            style="width: 80px; margin-left: 5px"
          >
            <el-option
              :key="count"
              v-for="(count, index) in payment"
              :label="count"
              :value="index + 1"
            ></el-option>
          </el-select>
          <el-select
            v-model="form.deposit_ratio"
            placeholder="请选择"
            style="width: 80px; margin-left: 5px"
          >
            <el-option
              :key="count"
              v-for="(count, index) in deposit"
              :label="count"
              :value="index + 1"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="房号配置">
          <el-button icon="el-icon-plus" size="mini" @click="addConfig"
            >添加</el-button
          >
        </el-form-item>
        <el-form-item v-if="form.house_facility_list.length !== 0">
          <template v-for="(config, index) in form.house_facility_list">
            <el-input
              :key="index"
              disabled
              style="width: 100px; margin-left: 5px; margin-top: 5px"
              :placeholder="config"
            ></el-input>
          </template>
        </el-form-item>
        <el-form-item label="房号照片">
          <el-upload
            class="avatar-uploader"
            action="https://jsonplaceholder.typicode.com/posts/"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
          >
            <img v-if="form.indexImg" :src="form.indexImg" class="avatar" />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i> </el-upload
        ></el-form-item>
        <div v-if="letVisible === true">
          <el-form-item label="最早入住日">
            <el-date-picker
              v-model="form.ruzhuri"
              align="right"
              type="date"
              placeholder="选择日期"
            >
            </el-date-picker>
          </el-form-item>
          <el-form-item label="房源描述">
            <el-input
              v-model="form.home_desc"
              clearable
              type="textarea"
              style="width: 400px"
              placeholder="请输入一些对房源的描述信息(如房屋亮点、招聘要求等)"
            ></el-input>
          </el-form-item>
          <el-form-item label="联系人">
            <el-radio-group size="small">
              <el-radio-button label="0">先生</el-radio-button>
              <el-radio-button label="1">女士</el-radio-button>
            </el-radio-group>
            <el-input
              v-model="form.house_contact_info.real_name"
              placeholder="请输入姓名"
              style="width: 200px; margin-left: 5px"
            ></el-input>
          </el-form-item>
          <el-form-item label="电话号码">
            <el-input
              style="width: 300px"
              v-model="form.house_contact_info.mobile"
            ></el-input>
          </el-form-item>
        </div>

        <el-form-item>
          <el-button v-if="letVisible === false" @click="back" size="medium"
            >上一步</el-button
          >
          <el-button v-else @click="cancel" size="medium">取消</el-button>
          <el-button
            v-if="letVisible === false"
            @click="saveAndLet"
            type="primary"
            size="medium"
            >保存并招租</el-button
          >
          <el-button v-else @click="publish" size="medium" type="primary"
            >发布招租</el-button
          ><el-button
            v-if="letVisible === false"
            @click="save"
            type="primary"
            size="medium"
            >保存</el-button
          >
        </el-form-item>
      </el-form>
    </div>

    <el-dialog title="选择收款账号" :visible.sync="dialogPayVisible">
      <div style="width: 100%; border: 1px red solid">
        <el-checkbox></el-checkbox>
        <el-button circle>
          <img src="@/assets/other_images/wechat.png" />
        </el-button>
      </div>
    </el-dialog>

    <el-dialog
      width="800px"
      title="请在地图上点击需要定位的位置(鼠标可拖动、放大、缩小地图)"
      :visible.sync="dialogMapVisible"
    >
      <div slot="footer" style="margin-top: -100px">
        <mapComponent
          @map-confirm="confirmMap"
          @cancel="closeMap"
        ></mapComponent>
      </div>
    </el-dialog>

    <el-dialog
      title="选择配置"
      :visible.sync="dialogConfigVisible"
      width="900px"
    >
      <el-form :model="form.config">
        <el-form-item>
          <el-checkbox-group v-model="config">
            <el-checkbox label="空调" name="type">空调</el-checkbox>
            <el-checkbox label="冰箱" name="type">冰箱</el-checkbox>
            <el-checkbox label="衣柜" name="type">衣柜</el-checkbox>
            <el-checkbox label="书桌" name="type">书桌</el-checkbox>
            <el-checkbox label="椅凳" name="type">椅凳</el-checkbox>
            <el-checkbox label="洗衣机" name="type">洗衣机</el-checkbox>
            <el-checkbox label="油烟机" name="type">油烟机</el-checkbox>
            <el-checkbox label="宽带" name="type">宽带</el-checkbox>
            <el-checkbox label="热水器" name="type">热水器</el-checkbox>
            <el-checkbox label="电视机" name="type">电视机</el-checkbox>
          </el-checkbox-group></el-form-item
        >
        <el-form-item
          >自定义配置
          <div style="float: right">
            <el-button
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="innerVisible = true"
              >添加</el-button
            >
            <el-button type="primary" icon="el-icon-setting" size="mini"
              >管理</el-button
            >
          </div>
        </el-form-item>
        <div v-if="otherConfig.length !== 0">
          <el-form-item>
            <el-checkbox-group v-model="checkOtherConfig">
              <el-checkbox
                style="display: block"
                v-for="config in otherConfig"
                :label="config"
                @change="otherConfigChange"
                :key="config"
                >{{ config }}</el-checkbox
              >
            </el-checkbox-group>
          </el-form-item>
        </div>

        <el-form-item style="text-align: center">
          <el-button type="primary" size="medium" @click="saveConfig"
            >确定</el-button
          >
        </el-form-item>
      </el-form>
      <el-dialog
        width="30%"
        title="添加自定义配置"
        :visible.sync="innerVisible"
        append-to-body
      >
        <el-form>
          <el-form-item label="请输入配置名称">
            <el-input v-model="otherConfigName"></el-input>
          </el-form-item>
          <el-form-item style="float: right">
            <el-button>取消</el-button>
            <el-button type="primary" @click="saveOtherConfig">确定</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </el-dialog>
  </div>
</template>

<script>
import mapComponent from "@/views/house/myMap.vue";
import utils from "@/api/utils";

export default {
  data() {
    return {
      form: {
        house_facility_list: [],
        house_contact_info: {
          real_name: "",
          mobile: "",
        },
      },
      dialogPayVisible: false, // 控制对话框是否显示
      dialogMapVisible: false, // 控制地图是否显示
      dialogConfigVisible: false, // 控制房号配置是否显示
      innerVisible: false, // 控制内层Dialog是否显示
      step: 1, // 步骤
      letVisible: false, // 控制招租信息是否显示
      floor: [],
      floorTotal: [],
      room: [],
      livingRoom: [],
      restRoom: [],
      payment: [],
      deposit: [],
      config: [], // 房号配置
      otherConfigName: "", // 房号配置中的自定义配置(配置名称)
      otherConfig: [], // 房号配置中的自定义配置(配置数组)
      checkOtherConfig: [], // 被选中的自定义配置
      province: "", // 地图定位，省
      city: "", // 地图定位，市
      district: "", // 地图定位，县/区
      street: "", // 地图定位，街道
      streetNumber: "", // 地图定位，门牌号
      imageUrl: "", // 房源展示图url地址
    };
  },
  components: {
    mapComponent,
  },
  mounted() {
    // 准备楼层数据
    this.readyFloorData();
    // 准备户型数据
    this.readyRoomData();
    // 准备收租周期数据
    this.readyRentData();
    // 接收房源列表传递过来的数据
    this.form = this.$route.query.form;
    this.imageUrl = this.form.index_img;
    // 处理省、市、县等地址信息
    // this.province = this.form
    // 处理房源设施信息
    var tempList = [];
    for (var i = 0; i < this.form.house_facility_list.length; ++i) {
      tempList.push(this.form.house_facility_list[i].name);
    }
    this.form.house_facility_list = tempList;
    // 处理电梯、楼梯信息
    if (this.form.has_elevator) {
      this.form.has_elevator = 0;
    } else {
      this.form.has_elevator = 1;
    }
  },
  methods: {
    // 准备楼层数据
    readyFloorData() {
      var floorTemp = [];
      var floorTotalTemp = [];
      for (var i = 1; i < 100; ++i) {
        floorTemp.push("第" + i + "层");
        floorTotalTemp.push("共" + i + "层");
      }
      this.floor = floorTemp;
      this.floorTemp = floorTotalTemp;
    },
    // 准备户型数据
    readyRoomData() {
      var roomTemp = [];
      var livingRoomTemp = [];
      var restRoomTemp = [];
      for (var i = 1; i < 10; ++i) {
        roomTemp.push(i + " 室");
        livingRoomTemp.push(i + " 厅");
        restRoomTemp.push(i + " 卫");
      }
      this.room = roomTemp;
      this.livingRoom = livingRoomTemp;
      this.restRoom = restRoomTemp;
    },
    // 准备收租周期数据
    readyRentData() {
      var paymentTemp = [];
      var depositTemp = [];
      for (var i = 1; i <= 12; ++i) {
        paymentTemp.push("付 " + i);
        depositTemp.push("压 " + i);
      }
      this.payment = paymentTemp;
      this.deposit = depositTemp;
    },
    // 弹出选择收款账号的对话框
    selectPayAccount() {
      this.dialogPayVisible = true;
    },
    // 取消跳转上一页
    cancel() {
      this.$router.push("/house/index");
    },
    // 显示地图
    showMap() {
      this.dialogMapVisible = true;
    },
    // 关闭地图
    closeMap() {
      this.dialogMapVisible = false;
    },
    // 获取子组件(地图)传递过来的经纬度
    confirmMap(center) {
      //用所定位的经纬度查找所在地省市街道等信息
      var point = new BMap.Point(center.lng, center.lat);
      var gc = new BMap.Geocoder();
      let _this = this;
      const myVue = this;
      gc.getLocation(point, function (rs) {
        var addComp = rs.addressComponents;
        // 自动填入
        myVue.form.address = rs.address;
        myVue.province = addComp.province;
        myVue.city = addComp.city;
        myVue.district = addComp.district;
        myVue.street = addComp.street;
        myVue.streetNumber = addComp.streetNumber;
      });
      // 关闭地图
      this.dialogMapVisible = false;
    },
    // 下一步
    next() {
      this.step = 2;
    },
    // 上一步
    back() {
      this.step = 1;
      this.letVisible = false;
    },
    // 保存并招租
    saveAndLet() {
      this.letVisible = true;
      // 保存房源数据
      this.saveHouseData();
    },
    // 保存
    save() {
      // 跳转回房源列表
      this.$router.push("/house/index");
      this.letVisible = false;
      // 保存房源数据
      this.saveHouseData();
    },
    // 房源展示图上传成功后调用此方法
    handleHouseImgSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
      this.form.index_img = res.data.file_key;
    },
    // 保存房源数据
    saveHouseData() {
      // 在此处发送请求保存房源信息，此时房源尚未开始招租
      // console.log(this.form);
    },
    // 发布招租
    publish() {
      // 跳转回房源列表
      this.$router.push("/house/index");
      // 先获取房源设施id
      utils.getFacilities().then((response) => {
        // 转换房源设施数据
        var facilityList = response.data.house_facility_list;
        var tempList = [];
        for (var i = 0; i < this.form.house_facility_list.length; ++i) {
          for (var j = 0; j < facilityList.length; ++j) {
            if (this.form.house_facility_list[i] === facilityList[j].name) {
              tempList.push({
                facility_id: facilityList[i].facility_id,
                name: facilityList[i].name,
                icon: facilityList[i].icon,
              });
            }
          }
        }
        this.form.house_facility_list = tempList;
        // 处理地址信息
        this.form.city = this.province + this.city;
        this.form.district = this.district;
        this.address =
          this.province +
          this.city +
          this.district +
          this.street +
          this.streetNumber;
        utils.publishHouse(this.form).then((response) => {});
      });
    },
    // 添加房号配置(家电)
    addConfig() {
      this.dialogConfigVisible = true;
    },
    // 保存房号配置
    saveConfig() {
      // 在这之前，将所有自定义配置从房号配置中移除掉
      for (var i = 0; i < this.otherConfig.length; ++i) {
        for (var j = 0; j < this.config.length; ++j) {
          if (this.otherConfig[i] === this.config[j]) {
            this.config.splice(j, 1);
          }
        }
      }

      // 判断是否勾选了自定义配置
      if (this.checkOtherConfig.length !== 0) {
        // 将自定义配置加入房号配置中
        for (var i = 0; i < this.checkOtherConfig.length; ++i) {
          this.config.push(this.checkOtherConfig[i]);
        }
      }

      // 将配置存入form中
      this.form.house_facility_list = this.config;

      // 关闭选择配置对话框
      this.dialogConfigVisible = false;

      // 将自定义配置保存到数据库
      for (var i = 0; i < this.otherConfig.length; ++i) {
        utils
          .publishFacility({
            name: this.otherConfig[i],
            icon: "",
          })
          .then((response) => {});
      }
    },
    // 保存房号配置中的自定义配置
    saveOtherConfig() {
      this.otherConfig.push(this.otherConfigName);
      this.innerVisible = false;
    },
    // 当自定义配置被选中时调用此函数
    otherConfigChange(value) {},
    // 当计数器改变时调用此函数
    handleChange() {},
    // 当房号图片上传时调用此函数
    handleAvatarSuccess() {},
    // 在房号图片上传之前调用此函数
    beforeAvatarUpload() {},
  },
};
</script>

<style scope>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
