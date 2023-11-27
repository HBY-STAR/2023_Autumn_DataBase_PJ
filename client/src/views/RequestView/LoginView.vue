<template>
  <div class="wrapper">
    <div
      style="
        margin: 180px auto;
        background-color: #fff;
        width: 300px;
        height: 280px;
        padding: 20px;
        border-radius: 10px;
      "
    >
      <div style="margin: 10px 0; text-align: center; font-size: 24px; color: cornflowerblue"><b>登录</b></div>
      <div style="height: 20px"></div>
      <el-form :model="user_login" :rules="rules" ref="userForm">
        <el-form-item prop="username">
          <el-input
            size="large"
            prefix-icon="User"
            v-model="user_login.username"
            placeholder="请输入用户名"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            size="large"
            prefix-icon="Lock"
            show-password
            v-model="user_login.password"
            placeholder="请输入密码"
          ></el-input>
        </el-form-item>
        <el-form-item prop="type">
          <el-radio-group v-model="user_login.type">
            <el-radio :label= "'user'" >用户</el-radio>
            <el-radio :label= "'seller'" >商家</el-radio>
            <el-radio :label= "'admin'" >管理员</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="default" autocomplete="off" @click="login">登录</el-button>
          <el-button autocomplete="off" style="position: absolute;right: 0" type="warning" @click="router().push('/')">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import router from "../../router";

export default {
  name: "user_login",
  data() {
    return {
      user_login: {
        username: null,
        password: null,
        type: null,
      },
      rules: {
        username: [
          { required: true, message: "用户名不能为空", trigger: "blur" },
          { min: 2, max: 20, message: "长度在 2 到 20 个字符", trigger: "blur" }
        ],
        password: [
          { required: true, message: "密码不能为空", trigger: "blur" },
          { min: 5, max: 20, message: "长度在 5 到 20 个字符", trigger: "blur" }
        ],
        type: [
          {required: true, trigger: "blur"}
        ]
      },
    };
  },
  methods: {
    router() {
      return router
    },
    login() {
      this.$refs["userForm"].validate((valid) => {
        if (valid) {
          // 表单校验合法
          this.request.post("/login", this.user).then((res) => {
            if (res.code === "200")
            {
              //存储token
              localStorage.setItem("token", JSON.stringify(res.data.access));
              if (this.user_login.type === "admin") {
                this.$router.push("/admin");
              } else if (this.user_login.type === "user") {
                this.$router.push("/user");
              } else if (this.user_login.type === "seller") {
                this.$router.push("/seller");
              } else {
                this.$router.push("/");
              }
              this.$message.success(res.data.message);
            }
            else
            {
              this.$message.error(res.data.message);
            }
          });
        }
      });
    }
  }
};
</script>

<style scoped>
.wrapper {
  height: 100vh;
  background-color: #909399;
  overflow: hidden;
}
</style>

