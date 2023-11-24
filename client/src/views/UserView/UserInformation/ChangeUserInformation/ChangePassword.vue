<template>
  <div style="margin-top: 10px">
    <el-form :model="user" :rules="rules" ref="userForm">
      <el-form-item prop="password">
        <el-input
            v-model="user.password"
            placeholder="请输入密码"
        ></el-input>
      </el-form-item>
    </el-form>
    <div style="margin-top: 30px">
      <el-button autocomplete="off" type="primary" @click="changePassword"
      >保存
      </el-button
      >
      <el-button
          autocomplete="off"
          style="position: absolute;right: 0;"
          type="warning"
          @click="$router.push('/personal_information')"
      >取消
      </el-button
      >
    </div>
  </div>
</template>

<script>

export default {
  name: "change_password",
  data() {
    return {
      user:{
        password: '',
        uid: 0,
      },
      rules:{
        password: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          { min: 6, max: 32, message: '长度在 6 到 32 个字符', trigger: 'blur' },
          { pattern: /^((?=.*\d)(?=.*[a-zA-Z]))|((?=.*[-_])(?=.*\d))|((?=.*[a-zA-Z])(?=.*[-_])).{6,32}$/,
            message: '字⺟，数字或者特殊字符（-_）⾄少包含两种'}
        ],
      },
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
    }
  },
  methods: {
    changePassword() {
      this.$refs['userForm'].validate((valid) => {
        if (valid) {
          this.user.uid = this.user_data.uid;
          this.request.post('/user-data/change_password', this.user).then((res) => {
            if (res.code === '200') {
              localStorage.setItem('user_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
              this.$message.success("保存成功")
            } else {
              this.$message.error(res.msg)
            }
          })
        }
      })
    },
  }
}
</script>