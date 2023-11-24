<template>
  <el-container>
    <el-header style="background-color: darkgrey;border-radius: 5px" height="65px">
      <el-container>
        <el-aside
          width="1330px"
          style="text-align: left; font-size: 20px; font-weight: 1000; color: darkslateblue"
        >
          <div  style="margin-top: 20px">
            <span>Group22_lab</span>
          </div>
        </el-aside>
        <el-main style="text-align: right; font-size: 16px">
          <div>
            <div style="width: 120px; float: left;margin-top: 2px">
              <el-avatar
                  style="float: left;margin-right: 10px;height: 30px;width: 30px;"
                  @click="upload_avatar_dialog=true"
                  :src=this.user_data.avatarUrl
              >
              </el-avatar>
              <el-dialog
                  v-model="this.upload_avatar_dialog"
                  title="上传头像"
                  width="20%"
                  style="text-align: center;"
                  fit="fit"
              >
                <div>
                  <el-upload
                      ref="upload"
                      drag
                      multiple
                      list-type="picture"
                      enctype="multipart/form-data"
                      action="#"
                      :limit="1"
                      :show-file-list="true"
                      :auto-upload="false"
                      :http-request="picUpload"
                      :before-upload="beforeAvatarUpload"
                  >
                    <img :src=this.user_data.avatarUrl
                         alt="头像"
                         style="height: 200px;width: 200px;margin: 20px"
                    >
                    <div class="el-upload__text">
                      将文件拖至此处，或者 <em>点击上传</em>
                    </div>
                    <template #tip>
                      <div class="el-upload__tip">
                        文件格式为jpg或png，且不大于5MB。
                      </div>
                    </template>
                  </el-upload>
                </div>
                <template #footer>
                    <span class="dialog-footer">
                      <el-button @click="this.upload_avatar_dialog = false" style="position: absolute;left: 30px">返回</el-button>
                      <el-button type="primary" @click="submitUpload">
                        确认
                      </el-button>
                    </span>
                </template>
              </el-dialog>
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
                    <message /> </el-icon
                  >
                  <div @click="router.push('/user/user_product_selling')">在售商品</div>
                </template>
              </el-menu-item>
              <el-menu-item index="2">
                <template #title>
                  <el-icon>
                    <icon-menu /> </el-icon
                  >
                  <div @click="router.push('/user/user_shop_open')">在营商店</div>
                </template>
              </el-menu-item>
              <el-menu-item index="3">
                <template #title>
                  <el-icon>
                    <setting /> </el-icon
                  >
                  <div @click="router.push('/seller/seller_shop_cart')">购物车</div>
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
  name: 'user_view',
  data() {
    return {
      user_data: localStorage.getItem('user_data')
        ? JSON.parse(localStorage.getItem('user_data'))
        : {},
      upload_avatar_dialog:false,
    }
  },
  created() {
    this.$router.push('/user/user_product_selling')
    this.findAvatar()
  },
  methods: {
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 5;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 5MB!');
      }
      return isJPG && isLt2M;
    },
//这里是自定义发送请求
    picUpload(f) {
      let params = new FormData()
      //注意在这里一个坑f.file
      params.append("file", f.file);
      this.$axios({
        method:'post',
        //这里的id是我要改变用户的ID值
        url:'/user-data/upload_image_'+this.user_data.uid,
        data:params,
        headers:{
          'content-type':'multipart/form-data'
        }
      }).then(res => {
        //这里是接受修改完用户头像后的JSON数据
        this.user_data.avatarUrl=res
        this.$message.success("上传成功")
      })
    },
    //触发请求
    submitUpload() {
      this.$refs.upload.submit();
    },
    findAvatar(){
      this.request.get('/user-data/find_'+this.user_data.uid).then((res)=>{
        localStorage.setItem('user_data', JSON.stringify(res))
      })
    },
  }
}
</script>

<script setup>
import { Menu as IconMenu, Message, Setting } from '@element-plus/icons-vue'
import { RouterLink } from 'vue-router'
import router from "@/router";
</script>
