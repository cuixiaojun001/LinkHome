import request from '@/utils/request'

export default {
    // -------- 用户管理模块 --------
    // 查询用户列表
    getUserList(param) {
        return request({
          url: `/api/v1/admin/user/users`,
          method: 'post',
          data:param
        })
    },
    // 根据用户id查询用户的详细信息
    getUserDetailById(id) {
        return request({
            url: `/api/v1/user/profile/${id}`,
            method: "get"
        })
    },
    // 根据用户id更新该用户的详细信息
    updateUserDetailById(id, userDetail){
        return request({
            url: `/api/v1/user/profile/${id}`,
            method: "put",
            data:userDetail
        })
     },

    // -------- 房源管理模块 --------
    // 获取房源列表
    getHouseList(param) {
        return request({
            url: `/api/v1/house/houses`,
            method: "post",
            data:param
        })
    },
    // 根据房源id下架房源
    dropHouseById(id,house) {
        return request({
            url: `/api/v1/admin/houses/${id}`,
            method: "put",
            data:house
        })
    },
    // 根据id查询指定房源详情信息
    getHouseDetailById(id) { 
        return request({
            url: `/api/v1/house/houses/${id}`,
            method: "get"
        })
    },
    // 发布房源
    publishHouse(houseInfo) { 
        return request({
            url: '/api/v1/house/publish',
            method: 'post',
            data:houseInfo
        })
    },
    // 获取房源设施
    getFacilities() { 
        return request({
            url: "/api/v1/house/facilities",
            method:"get"
        })
    },
    // 发布房源设施
    publishFacility(facility) { 
        return request({
            url: "/api/v1/house/facilities/",
            method: "post",
            data:facility
        })
    },

    // -------- 订单管理模块 --------
    // 获取订单列表
    getOrderList(param) { 
        return request({
            url: `/api/v1/admin/orders`,
            method:"post",
            data:param
        })
    }
}