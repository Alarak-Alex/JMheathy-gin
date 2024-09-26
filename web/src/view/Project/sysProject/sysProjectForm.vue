<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标题列表:" prop="titleList">
          <ArrayCtrl v-model="formData.titleList" editable />
        </el-form-item>
        <el-form-item label="图片类型:" prop="picType">
          <el-select v-model="formData.picType" placeholder="请选择图片类型" style="width:100%" :clearable="true">
            <el-option v-for="(item, key) in PicTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="提示词ID:" prop="promtId">
          <el-input v-model.number="formData.promtId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="cookie类型:" prop="cookieType">
          <el-select v-model="formData.cookieType" placeholder="请选择cookie类型" style="width:100%" :clearable="true">
            <el-option v-for="(item, key) in CookieTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="管理用户ID:" prop="systemUserId">
          <el-input v-model.number="formData.systemUserId" :clearable="true" placeholder="请输入" />
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
  createSystemProject,
  updateSystemProject,
  findSystemProject
} from '@/api/Project/sysProject'

defineOptions({
  name: 'SystemProjectForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl2.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const PicTypeOptions = ref([])
const CookieTypeOptions = ref([])
const formData = ref({
  titleList: [],
  picType: '',
  promtId: undefined,
  cookieType: '',
  systemUserId: undefined,
})
// 验证规则
const rule = reactive({
  titleList: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  picType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  promtId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  cookieType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  systemUserId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findSystemProject({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  PicTypeOptions.value = await getDictFunc('PicType')
  CookieTypeOptions.value = await getDictFunc('CookieType')
}

init()
// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSystemProject(formData.value)
        break
      case 'update':
        res = await updateSystemProject(formData.value)
        break
      default:
        res = await createSystemProject(formData.value)
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

<style></style>
