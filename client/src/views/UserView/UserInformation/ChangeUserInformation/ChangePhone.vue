<template>
  <div style="margin-top: 10px">
    <el-form :model="user" :rules="rules" ref="userForm">
      <el-form-item prop="phone">
        <el-input
            v-model="user.phone"
            placeholder="请输入手机号码"
        ></el-input>
      </el-form-item>
    </el-form>
    <div style="margin-top: 30px">
      <el-button autocomplete="off" type="primary" @click="changePhone"
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
  name: "change_phone",
  data() {
    return {
      user:{
        phone: '',
        uid: 0,
      },
      rules:{
        phone: [
          { required: true, message: '手机号不能为空', trigger: 'blur' },
          { min: 11, max: 11, message: '长度应为 11 个数字', trigger: 'blur' },
          { pattern: /^1[34578]\d{9}$/, message: "请输入正确的手机号（目前只支持中国大陆）", trigger: "blur",},
        ],
      },
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
    }
  },
  methods: {
    changePhone() {
      this.$refs['userForm'].validate((valid) => {
        if (valid) {
          this.user.uid = this.user_data.uid;
          this.request.post('/user-data/change_phone', this.user).then((res) => {
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