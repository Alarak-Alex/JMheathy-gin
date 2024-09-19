<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>

        <!-- <el-form-item label="用户Cookie" prop="userCookie">
          <el-input v-model="searchInfo.userCookie" placeholder="搜索条件" />

        </el-form-item> -->
        <el-form-item label="Cookie类型" prop="cookieType">
          <el-select v-model="searchInfo.cookieType" clearable placeholder="请选择"
            @clear="() => { searchInfo.cookieType = undefined }">
            <el-option v-for="(item, key) in CookieTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="管理用户ID" prop="systemUserId">

          <el-input v-model.number="searchInfo.systemUserId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="Cookie名" prop="cookieName">
          <el-input v-model="searchInfo.cookieName" placeholder="搜索条件" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery = true"
            v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery = false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;"
          :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="Cookie_Cookie" />
        <ExportExcel v-auth="btnAuth.exportExcel" template-id="Cookie_Cookie" />
        <ImportExcel v-auth="btnAuth.importExcel" template-id="Cookie_Cookie" @on-success="getTableData" />
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange" @sort-change="sortChange">
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <!-- <el-table-column label="用户Cookie" prop="userCookie" width="200">
          <template #default="scope">
            <div class="file-list">
              <el-tag v-for="file in scope.row.userCookie" :key="file.uid" @click="downloadFile(file.url)">{{ file.name
                }}</el-tag>
            </div>
          </template>
        </el-table-column> -->
        <el-table-column sortable align="left" label="Cookie类型" prop="cookieType" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.cookieType, CookieTypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="是否启用" prop="enable" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.enable, BoolTypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="管理用户ID" prop="systemUserId" width="120" />
        <el-table-column sortable align="left" label="Cookie名" prop="cookieName" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button"
              @click="getDetails(scope.row)"><el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button"
              @click="updateCookieFunc(scope.row)">变更</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete"
              @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="用户Cookie:" prop="userCookie">
          <!-- <SelectFile v-model="formData.userCookie" /> -->
          <el-form-item label="上传Yaml:" prop="titleList">
            <!-- <el-upload class="upload" :action="`${getBaseUrl()}/fileUploadAndDownload/upload?noSave=1`"
              :on-change="onChange" :show-file-list="true">
              <el-button type="primary">点击上传Yaml文件</el-button>
              <template #tip>
                <div class="el-upload__tip">
                  仅支持上传 .yaml 文件
                </div>
              </template>
            </el-upload> -->
            <UploadYaml v-model="formData.userCookie" />
          </el-form-item>
        </el-form-item>
        <el-form-item label="Cookie类型:" prop="cookieType">
          <el-select v-model="formData.cookieType" placeholder="请选择Cookie类型" style="width:100%" :clearable="true">
            <el-option v-for="(item, key) in CookieTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否启用:" prop="enable">
          <el-select v-model="formData.enable" placeholder="请选择是否启用" style="width:100%" :clearable="true">
            <el-option v-for="(item, key) in BoolTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="管理用户ID:" prop="systemUserId">
          <el-input v-model.number="formData.systemUserId" :clearable="true" placeholder="请输入管理用户ID" />
        </el-form-item>
        <el-form-item label="Cookie名:" prop="cookieName">
          <el-input v-model="formData.cookieName" :clearable="true" placeholder="请输入Cookie名" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <!-- <el-descriptions-item label="用户Cookie">
          <div class="fileBtn" v-for="(item, index) in detailFrom.userCookie" :key="index">
            <el-button type="primary" text bg @click="onDownloadFile(item.url)">
              <el-icon style="margin-right: 5px">
                <Download />
              </el-icon>
              {{ item.name }}
            </el-button>
          </div>
        </el-descriptions-item> -->
        <el-descriptions-item label="Cookie类型">
          {{ detailFrom.cookieType }}
        </el-descriptions-item>
        <el-descriptions-item label="是否启用">
          {{ detailFrom.enable }}
        </el-descriptions-item>
        <el-descriptions-item label="管理用户ID">
          {{ detailFrom.systemUserId }}
        </el-descriptions-item>
        <el-descriptions-item label="Cookie名">
          {{ detailFrom.cookieName }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createCookie,
  deleteCookie,
  deleteCookieByIds,
  updateCookie,
  findCookie,
  getCookieList
} from '@/api/Cookie/sysCookie'
import { getUrl } from '@/utils/image'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'
import UploadYaml from '@/components/uploadyaml/uploadyaml.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { getBaseUrl } from '@/utils/format'
import * as jsyaml from 'js-yaml';
// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
  name: 'Cookie'
})
// 按钮权限实例化
const btnAuth = useBtnAuth()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
  userCookie: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  cookieType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  enable: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  cookieName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    userCookie: 'user_cookie',
    cookieType: 'cookie_type',
    enable: 'enable',
    systemUserId: 'system_user_id',
    cookieName: 'cookie_name',
  }

  let sort = sortMap[prop]
  if (!sort) {
    sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getCookieList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  CookieTypeOptions.value = await getDictFunc('CookieType')
  BoolTypeOptions.value = await getDictFunc('BoolType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteCookieFunc(row)
  })
}

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
    const res = await deleteCookieByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCookieFunc = async (row) => {
  const res = await findCookie({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteCookieFunc = async (row) => {
  const res = await deleteCookie({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 文件选择回调
const onChange = (file, fileList) => {
  ReadYaml(file.raw);
};

// 读取 YAML 文件
const ReadYaml = (file) => {
  const reader = new FileReader();
  reader.onload = (event) => {
    try {
      const yamlContent = event.target.result;
      const jsonObject = jsyaml.load(yamlContent);
      formData.value.userCookie = JSON.stringify(jsonObject, null, 2);
      console.log(formData.value.userCookie);
    } catch (error) {
      console.error('Error in ReadYaml:', error);
      ElMessage({ type: 'error', message: error.message });
    }
  };
  reader.readAsText(file);
};

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    userCookie: [],
    cookieType: '',
    enable: '',
    systemUserId: undefined,
    cookieName: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
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
      closeDialog()
      getTableData()
    }
  })
}

const downloadFile = (url) => {
  window.open(getUrl(url), '_blank')
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCookie({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>
.file-list {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn {
  margin-bottom: 10px;
}

.fileBtn:last-child {
  margin-bottom: 0;
}
</style>