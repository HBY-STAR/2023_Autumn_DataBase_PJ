<template>
  <el-container>
    <el-header style="background-color: darkgrey;border-radius: 5px" height="65px">
      <el-container>
        <el-aside
          width="500px"
          style="text-align: left; font-size: 20px; font-weight: 1000; color: darkslateblue"
        >
          <div style="margin-top: 15px">
            <span>PriceCompareSystem</span>
          </div>
        </el-aside>
        <el-main style="text-align: right; font-size: 16px">
          <div>
            <span style="position: relative">{{IsAdmin()}}</span>
            <el-dropdown>
              <el-icon style="margin-left: 8px; margin-top: 1px; font-size: large">
                <setting />
              </el-icon>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>
                    <RouterLink to="/login">切换账号</RouterLink>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <RouterLink to="/" @click="localStorage.removeItem('token') ">退出登录</RouterLink>
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </el-main>
      </el-container>
    </el-header>
    <el-container>
      <el-aside width="200px" height="640px">
        <el-container>
          <el-header
            style="text-align: center; font-size: 18px; background-color: antiquewhite;border-radius: 5px;margin-right: 20px"
            height="50px"
          >
            <div style="margin-top: 10px;">
              <span>管理员界面</span>
            </div>
          </el-header>
          <el-main>
            <el-menu>
              <el-menu-item index="1">
                <template #title>
                  <el-icon>
                    <user />
                  </el-icon>
                  <div @click="router.push('/admin/all_user_data')">用户管理</div>
                </template>
              </el-menu-item>
              <el-menu-item index="2">
                <template #title>
                  <el-icon>
                    <house />
                  </el-icon>
                  <div @click="router.push('/admin/all_seller_data')">商家管理</div>
                </template>
              </el-menu-item>
              <el-menu-item index="3">
                <template #title>
                  <el-icon>
                    <place />
                  </el-icon>
                  <div @click="router.push('/admin/all_platform_data')">平台管理</div>
                </template>
              </el-menu-item>
              <el-menu-item index="4">
                <template #title>
                  <el-icon>
                    <sell />
                  </el-icon>
                  <div @click="router.push('/admin/all_admin_commodity_item_data')">已发布商品管理</div>
                </template>
              </el-menu-item>
              <el-menu-item index="5">
                <template #title>
                  <el-icon>
                    <goods />
                  </el-icon>
                  <div @click="router.push('/admin/all_commodity_data')">可发布商品管理</div>
                </template>
              </el-menu-item>
              <el-menu-item index="6">
                <template #title>
                  <el-icon>
                    <star />
                  </el-icon>
                  <div @click="router.push('/admin/statistics_favorite')">收藏统计</div>
                </template>
              </el-menu-item>
              <el-menu-item index="7">
                <template #title>
                  <el-icon>
                    <shopping-cart />
                  </el-icon>
                  <div @click="router.push('/admin/statistics_commodity_price')">价格统计</div>
                </template>
              </el-menu-item>
              <el-menu-item index="8">
                <template #title>
                  <el-icon>
                    <clock />
                  </el-icon>
                  <div @click="router.push('/admin/statistics_annual')">年度统计</div>
                </template>
              </el-menu-item>
            </el-menu>
          </el-main>
        </el-container>
      </el-aside>
      <el-main height="605px">
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>

export default {
  name: "admin_view",
  data() {
    return {
    };
  },
  created() {
  },
  methods: {
    IsAdmin(){
      return localStorage.getItem("user_type")  ? 'admin' : '未登录'
    }
  }
};

</script>

<script setup>

import { Clock, Money, Setting, ShoppingCart } from "@element-plus/icons-vue";
import { RouterLink } from "vue-router";
import router from "@/router";

</script>

