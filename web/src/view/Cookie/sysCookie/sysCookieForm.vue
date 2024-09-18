
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户Cookie:" prop="userCookie">
          <SelectFile v-model="formData.userCookie" />
       </el-form-item>
        <el-form-item label="Cookie类型:" prop="cookieType">
           <el-select v-model="formData.cookieType" placeholder="请选择Cookie类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in CookieTypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="是否启用:" prop="enable">
           <el-select v-model="formData.enable" placeholder="请选择是否启用" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in BoolTypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="管理用户ID:" prop="systemUserId">
          <el-input v-model.number="formData.systemUserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="Cookie名:" prop="cookieName">
          <el-input v-model="formData.cookieName" :clearable="true"  placeholder="请输入Cookie名" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCookie,
  updateCookie,
  findCookie
} from '@/api/Cookie/sysCookie'

defineOptions({
    name: 'CookieForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const CookieTypeOptions = ref([])
const BoolTypeOptions = ref([])
const formData = ref({
            userCookie: [],
            cookieType: '',
            enable: '',
            systemUserId: undefined,
            cookieName: '',
        })
// 验证规则
const rule = reactive({
               userCookie : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cookieType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               enable : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cookieName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCookie({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    CookieTypeOptions.value = await getDictFunc('CookieType')
    BoolTypeOptions.value = await getDictFunc('BoolType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCookie(formData.value)
               break
             case 'update':
               res = await updateCookie(formData.value)
               break
             default:
               res = await createCookie(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
