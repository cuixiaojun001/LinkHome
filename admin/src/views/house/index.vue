<template>
  <div>
    <div
      v-if="houseList.length === 0"
      style="position: absolute; top: 20%; left: 40%"
    >
      <img src="@/assets/other_images/no_house.png" />
      <p style="text-align: center; font-size: 22px">暂无房源</p>
      <p style="text-align: center">
        <el-button
          type="primary"
          size="medium"
          icon="el-icon-plus"
          @click="addHouse"
          >立刻添加</el-button
        >
        <el-button type="primary" size="medium" icon="el-icon-upload"
          >批量导入</el-button
        >
      </p>
    </div>
    <div v-if="houseList.length >= 0">
      <el-button
        type="primary"
        size="medium"
        style="margin: 10px"
        @click="addHouse"
        >添加房源</el-button
      >
      <el-table :data="houseList" border stripe>
        <el-table-column label="#" width="60" fixed>
          <template slot-scope="scope">
            {{ (page - 1) * limit + scope.$index + 1 }}
          </template>
        </el-table-column>
        <el-table-column
          label="房源标题"
          prop="title"
          align="center"
          width="300px"
        >
          <template slot-scope="scope">
            <img
              style="float: left"
              src="@/assets/other_images/house.png"
              width="45px"
            />
            <span style="float: left; margin-left: 5px; font-size: 15px">{{
              scope.row.title
            }}</span>
            <br />
            <el-link
              type="primary"
              style="margin-left: 5px; font-size: 10px"
              @click="addTenant"
              >添加租客</el-link
            >
            <el-link
              type="primary"
              style="margin-left: 5px; font-size: 10px"
              @click="editHouseInfo(scope.row.house_id)"
              >修改</el-link
            >
            <el-link type="primary" style="margin-left: 5px; font-size: 10px"
              >删除</el-link
            >
          </template>
        </el-table-column>
        <el-table-column
          label="房源地址"
          prop="address"
          align="center"
          width="300px"
        />
        <el-table-column label="租金" prop="rent_money" align="center">
          <template slot-scope="scope">
            <span>¥{{ scope.row.rent_money }}</span>
          </template>
        </el-table-column>
        <el-table-column label="出租类型" prop="rent_type" align="center">
          <template slot-scope="scope">
            <el-tag type="success" v-if="scope.row.rent_type === 'whole'"
              >整租</el-tag
            >
            <el-tag type="success" v-if="scope.row.rent_type === 'share'"
              >合租</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column label="出租状态" prop="rent_state" align="center">
          <template slot-scope="scope">
            <el-tag type="danger" v-if="scope.row.rent_state === 'rent'"
              >已出租</el-tag
            >
            <el-tag type="warning" v-if="scope.row.rent_state === 'not_rent'"
              >未出租</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column label="房屋类型" prop="house_type" align="center">
          <template slot-scope="scope">
            <el-tag type="primary" v-if="scope.row.house_type === 'department'"
              >公寓</el-tag
            >
            <el-tag type="primary" v-if="scope.row.house_type === 'community'"
              >小区</el-tag
            >
            <el-tag type="primary" v-if="scope.row.house_type === 'residential'"
              >普通住宅</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column label="审核状态" prop="state" align="center">
          <template slot-scope="scope">
            <el-tag type="primary" v-if="scope.row.state === 'up'"
              >已上架</el-tag
            >
            <el-tag type="info" v-if="scope.row.state === 'down'"
              >已下架</el-tag
            >
            <el-tag type="warning" v-if="scope.row.state === 'auditing'"
              >审核中</el-tag
            >
            <el-tag type="danger" v-if="scope.row.state === 'deleted'"
              >已删除</el-tag
            >
          </template> </el-table-column
        ><el-table-column label="操作" align="center" width="250" fixed="right">
          <template slot-scope="scope">
            <el-button
              type="primary"
              size="mini"
              @click="showHouseInfo(scope.row.house_id)"
            >
              查看
            </el-button>
            <el-button
              v-if="scope.row.state === 'auditing'"
              type="warning"
              size="mini"
              @click="toAuditDetail(scope.row.user_id, 1)"
            >
              审批
            </el-button>
            <el-button
              v-if="scope.row.state === 'up'"
              type="warning"
              size="mini"
              @click="drop(scope.row.house_id)"
            >
              下架
            </el-button>
            <el-button
              v-if="
                scope.row.state === 'down' || scope.row.state === 'auditing'
              "
              type="success"
              size="mini"
              @click="put(scope.row.house_id)"
            >
              上架
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-dialog :title="houseName" :visible.sync="dialogHouseDetailVisible">
        <table>
          <tr style="line-height: 50px">
            <td style="width: 150px">房号</td>
            <td style="width: 150px">{{ houseDetail.house_id }}</td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">楼层</td>
            <td>第{{ houseDetail.floor }}楼/共{{ houseDetail.max_floot }}楼</td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">户型</td>
            <td>
              {{ houseDetail.bedroom_num }}房{{
                houseDetail.living_room_num
              }}厅{{ houseDetail.toilet_num }}卫
            </td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">面积</td>
            <td>{{ houseDetail.area }}平米</td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">月租金</td>
            <td>¥{{ houseDetail.rent_money }}元</td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">收租周期</td>
            <td>付1压3</td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">照片</td>
            <td>
              <img
                :src="houseDetail.index_img"
                width="100px"
                @click="magnify(houseDetail.index_img)"
              />
            </td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">房间配置</td>
            <td>
              <el-input
                v-if="houseDetail.house_facility_list === []"
                disabled
                size="mini"
                style="width: 80px"
                >暂无房间配置</el-input
              >
              <div style="width: 500px">
                <el-button
                  type="info"
                  v-for="(facility, index) in houseDetail.house_facility_list"
                  :key="index"
                  size="mini"
                  disabled
                  >{{ facility.name }}</el-button
                >
              </div>
            </td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">房间亮点</td>
            <td>
              <el-input
                v-if="houseDetail.house_facility_list === []"
                disabled
                size="mini"
                style="width: 80px"
                >暂无房间配置</el-input
              >
              <div style="width: 500px">
                <el-button
                  type="info"
                  v-for="(facility, index) in houseDetail.house_facility_list"
                  :key="index"
                  size="mini"
                  disabled
                  >{{ facility.name }}</el-button
                >
              </div>
            </td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">招租要求</td>
            <td>
              <el-input
                v-if="houseDetail.house_facility_list === []"
                disabled
                size="mini"
                style="width: 80px"
                >暂无房间配置</el-input
              >
              <div style="width: 500px">
                <el-button
                  type="info"
                  v-for="(facility, index) in houseDetail.house_facility_list"
                  :key="index"
                  size="mini"
                  disabled
                  >{{ facility.name }}</el-button
                >
              </div>
            </td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">看房联系人</td>
            <td>{{ houseContactInfo.real_name }}</td>
          </tr>
          <tr style="line-height: 50px">
            <td style="width: 150px">联系电话</td>
            <td>{{ houseContactInfo.mobile }}</td>
          </tr>
        </table>
        <img v-if="innerVisible === true" :src="bigImageUrl" />
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
  </div>
</template>

<script>
import utilsApi from "@/api/utils.js";

export default {
  data() {
    return {
      houseList: [{ state: "auditing" }],
      total: 0, // 数据库中的总记录数
      page: 1, // 默认页码
      limit: 5, // 每页记录数
      dialogHouseDetailVisible: false, // 控制房源详情对话框是否显示
      houseName: "", // 房源详情对话框标题
      houseDetail: {}, // 房源详情
      houseContactInfo: {}, // 联系人信息
      innerVisible: false, // 控制放大照片的对话框是否显示
      bigImageUrl: "", // 放大的照片url
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    // 获取房源列表
    fetchData() {
      // 封装分页参数
      var param = {};
      param.offset = (this.page - 1) * this.limit;
      param.limit = this.limit;
      param.query_params = {};
      utilsApi.getHouseList(param).then((response) => {
        this.houseList = response.data.data_list;
        this.total = response.data.total;
      });
    },
    // 下架房源
    drop(id) {
      var tempHouse = {};
      for (var i = 0; i < this.houseList.length; ++i) {
        if (this.houseList[i].house_id === id) {
          tempHouse = this.houseList[i];
          // 将审核状态设置为已下架
          tempHouse.state = "down";
        }
      }
      utilsApi.dropHouseById(id, tempHouse).then((response) => {
        this.$message.success("下架成功");
      });
    },
    // 上架房源
    put(id) {
      var tempHouse = {};
      for (var i = 0; i < this.houseList.length; ++i) {
        if (this.houseList[i].house_id === id) {
          tempHouse = this.houseList[i];
          // 将审核状态设置为已上架
          tempHouse.state = "up";
        }
      }
      utilsApi.dropHouseById(id, tempHouse).then((response) => {
        this.$message.success("上架成功");
      });
    },
    // 查看房源详细信息
    showHouseInfo(id) {
      this.dialogHouseDetailVisible = true;
      utilsApi.getHouseDetailById(id).then((response) => {
        this.houseDetail = response.data;
        this.houseName = this.houseDetail.title;
        this.houseContactInfo = this.houseDetail.house_contact_info;
      });
    },
    // 修改房源信息
    editHouseInfo(id) {
      var form = {};
      // 获取房源信息进行回显
      utilsApi.getHouseDetailById(id).then((response) => {
        form = response.data;
        // 跳转至添加房源界面
        this.$router.push({
          path: "/house/add",
          query: {
            form: form,
          },
        });
      });
    },
    // 点击照片进行放大操作
    magnify(url) {
      this.innerVisible = true;
      this.bigImageUrl = url;
    },
    // 跳转至添加房源界面
    addHouse() {
      this.$router.push("/house/add");
    },
    // 跳转至添加租客界面
    addTenant() {
      this.$router.push("/house/addTenant");
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

<style></style>
