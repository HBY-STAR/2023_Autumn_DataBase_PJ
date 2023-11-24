<template>
  <div style="margin-top: 10px">
    <el-form :model="user" :rules="rules" ref="userForm">
      <el-form-item prop="email">
        <el-input
            v-model="user.email"
            placeholder="请输入邮箱"
        ></el-input>
      </el-form-item>
    </el-form>
    <div style="margin-top: 30px">
      <el-button autocomplete="off" type="primary" @click="changeEmail"
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
  name: "change_email",
  data() {
    return {
      user:{
        email: '',
        uid: 0,
      },
      rules:{
        email: [
          { required: true, message: '邮箱不能为空', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' },
        ],
      },
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
    }
  },
  methods: {
    changeEmail() {
      this.$refs['userForm'].validate((valid) => {
        if (valid) {
          this.user.uid = this.user_data.uid;
          this.request.post('/user-data/change_email', this.user).then((res) => {
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