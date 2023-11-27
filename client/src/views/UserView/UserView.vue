<template>
  <el-container>
    <el-header style="background-color: darkgrey;border-radius: 5px" height="65px">
      <el-container>
        <el-aside
          width="1330px"
          style="text-align: left; font-size: 20px; font-weight: 1000; color: darkslateblue"
        >
          <div style="margin-top: 20px">
            <span>Group22_lab</span>
          </div>
        </el-aside>
        <el-main style="text-align: right; font-size: 16px">
          <div>
            <div style="width: 120px; float: left;margin-top: 2px">
              <span style="float: left">{{ user_data.username }}</span>
              <el-dropdown style="float: left">
                <el-icon style="margin-left: 8px; margin-top: 1px; font-size: large">
                  <setting />
                </el-icon>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item>
                      <RouterLink to="/personal_information">个人信息</RouterLink>
                    </el-dropdown-item>
                    <el-dropdown-item>
                      <RouterLink to="/">退出登录</RouterLink>
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
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
            <div style="margin-top: 10px">
              <span>商品种类</span>
            </div>
          </el-header>
          <el-main>
            <el-menu>
              <el-menu-item index="1">
                <template #title>
                  <el-icon>
                    <message />
                  </el-icon
                  >
                  <div @click="router.push('/user/user_product_selling')">在售商品</div>
                </template>
              </el-menu-item>
              <el-menu-item index="2">
                <template #title>
                  <el-icon>
                    <icon-menu />
                  </el-icon
                  >
                  <div @click="router.push('/user/user_shop_open')">在营商店</div>
                </template>
              </el-menu-item>
              <el-menu-item index="3">
                <template #title>
                  <el-icon>
                    <setting />
                  </el-icon
                  >
                  <div @click="router.push('/seller/seller_shop_cart')">购物车</div>
                </template>
              </el-menu-item>
              <el-menu-item index="4">
                <template #title>
                  <el-icon>
                    <message />
                  </el-icon
                  >
                  To do
                </template>
              </el-menu-item>
              <el-menu-item index="5">
                <template #title>
                  <el-icon>
                    <message />
                  </el-icon
                  >
                  To do
                </template>
              </el-menu-item>
              <el-menu-item index="6">
                <template #title>
                  <el-icon>
                    <message />
                  </el-icon
                  >
                  To do
                </template>
              </el-menu-item>
            </el-menu>
          </el-main>
        </el-container>
      </el-aside>
      <el-main height="605px">
        <RouterView></RouterView>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  name: "user_view",
  data() {
    return {
      user_data: localStorage.getItem("user_data")
        ? JSON.parse(localStorage.getItem("user_data"))
        : {},
      home_commodity_all: localStorage.getItem("home_commodity_all")
        ? JSON.parse(localStorage.getItem("home_commodity_all"))
        : {},
    };
  },
  created() {
  },
  methods: {
  }
};
</script>

<script setup>
import { Menu as IconMenu, Message, Setting } from "@element-plus/icons-vue";
import { RouterLink } from "vue-router";
import router from "@/router";
</script>
