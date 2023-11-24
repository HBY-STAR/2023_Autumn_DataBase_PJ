<template>
  <el-container>
    <el-header style="background-color: darkgrey;border-radius: 5px" height="60px">
      <el-container>
        <el-aside
          width="500px"
          class="layout_header"
          style="text-align: left; font-size: 20px; font-weight: 1000; color: darkslateblue"
        >
          <div class="toolbar">
            <span>Lab2后台管理系统</span>
          </div>
        </el-aside>
        <el-main class="layout_header" style="text-align: right; font-size: 16px">
          <div class="toolbar">
            <span>sys_user</span>
            <el-dropdown>
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
        </el-main>
      </el-container>
    </el-header>
    <el-container>
      <el-aside width="200px" height="640px">
        <el-container>
          <el-header
            class="layout_header"
            style="text-align: center; font-size: 18px; background-color: antiquewhite;border-radius: 5px;margin-right: 20px"
            height="50px"
          >
            <div class="toolbar">
              <span>信息菜单</span>
            </div>
          </el-header>
          <el-main>
            <el-menu>
              <el-menu-item index="1">
                <template #title>
                  <el-icon> <message /> </el-icon>
                  <div @click="router.push('/sys_user/all_user_data')">用户数据</div>
                </template>
              </el-menu-item>
              <el-menu-item index="2">
                <template #title>
                  <el-icon> <message /> </el-icon>
                  <div @click="router.push('/sys_user/all_user_account')">用户账户</div>
                </template>
              </el-menu-item>
              <el-menu-item index="3">
                <template #title>
                  <el-icon> <icon-menu /> </el-icon>
                  <div @click="router.push('/sys_user/all_shop_data')">商店数据</div>
                </template>
              </el-menu-item>
              <el-menu-item index="4">
                <template #title>
                  <el-icon> <icon-menu /> </el-icon>
                  <div @click="router.push('/sys_user/all_shop_account')">商店账户</div>
                </template>
              </el-menu-item>
              <el-menu-item index="5">
                <template #title>
                  <el-icon> <icon-menu /> </el-icon>
                  <div @click="router.push('/sys_user/all_product_data')">商品数据</div>
                </template>
              </el-menu-item>
              <el-menu-item index="6">
                <template #title>
                  <el-icon> <setting /> </el-icon>
                  <div @click="router.push('/sys_user/shop_request_data')">商店开店请求</div>
                </template>
              </el-menu-item>
              <el-menu-item index="7">
                <template #title>
                  <el-icon> <setting /> </el-icon>
                  <div @click="router.push('/sys_user/shop_delete_data')">商店注销请求</div>
                </template>
              </el-menu-item>
              <el-menu-item index="8">
                <template #title>
                  <el-icon> <setting /> </el-icon>
                  <div @click="router.push('/sys_user/product_request_data')">商品上架请求</div>
                </template>
              </el-menu-item>
              <el-menu-item index="9">
                <template #title>
                  <el-icon> <setting /> </el-icon>
                  <div @click="router.push('/sys_user/not_checked_trade_record_data')">商城中间账户</div>
                </template>
              </el-menu-item>
              <el-menu-item index="10">
                <template #title>
                  <el-icon> <setting /> </el-icon>
                  <div @click="router.push('/sys_user/all_trade_record_data')">交易记录</div>
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
  name: 'sys_user_view',
  data() {
    return {
      form: {},
      user_data: localStorage.getItem('user_data')
        ? JSON.parse(localStorage.getItem('user_data'))
        : {}
    }
  }
}
</script>

<script setup>
import { Menu as IconMenu, Message, Setting } from '@element-plus/icons-vue'
import { RouterLink } from 'vue-router'
import router from "@/router";

</script>

<style>
.layout_header {
  position: relative;
}

.layout_header .toolbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
}
</style>
