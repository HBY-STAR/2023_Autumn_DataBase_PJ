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
            <span>未登录</span>
            <el-dropdown>
              <el-icon style="margin-left: 8px; margin-top: 1px; font-size: large">
                <setting />
              </el-icon>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>
                    <RouterLink to="/login">登录</RouterLink>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <RouterLink to="/register">注册</RouterLink>
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
            <div style="margin-top: 10px;">
              <span>商品种类</span>
            </div>
          </el-header>
          <el-main>
            <el-menu>
              <el-menu-item index="1">
                <template #title>
                  <el-icon>
                    <message /> </el-icon
                  >
                  <div @click="router.push('/home_product_selling')">在售商品</div>
                </template>
              </el-menu-item>
              <el-menu-item index="2">
                <template #title>
                  <el-icon>
                    <icon-menu /> </el-icon
                  >
                  <div @click="router.push('/home_shop_open')">在营商店</div>
                </template>
              </el-menu-item>
              <el-menu-item index="3">
                <template #title>
                  <el-icon>
                    <setting /> </el-icon
                  >To do
                </template>
              </el-menu-item>
              <el-menu-item index="4">
                <template #title>
                  <el-icon>
                    <message /> </el-icon
                  >To do
                </template>
              </el-menu-item>
              <el-menu-item index="5">
                <template #title>
                  <el-icon>
                    <message /> </el-icon
                  >To do
                </template>
              </el-menu-item>
              <el-menu-item index="6">
                <template #title>
                  <el-icon>
                    <message /> </el-icon
                  >To do
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
  name: 'home_view',
  data(){
    return{
      shop_open_data: localStorage.getItem('shop_open_data')
          ? JSON.parse(localStorage.getItem('shop_open_data'))
          : {}
    }
  },
  created() {
    // 从后台获取最新的User数据
    this.findAll()
    this.$router.push('/home_product_selling')

  },
  methods:{
    findAll(){
      this.request.get('/shop-data/find_open_shop').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('shop_open_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    }
  }
}
</script>

<script setup>
import { Menu as IconMenu, Message, Setting } from '@element-plus/icons-vue'
import { RouterLink } from 'vue-router'
import router from "@/router";

</script>

