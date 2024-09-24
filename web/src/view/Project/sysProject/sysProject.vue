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

        <el-form-item label="图片类型" prop="picType">
          <el-select v-model="searchInfo.picType" clearable placeholder="请选择"
            @clear="() => { searchInfo.picType = undefined }">
            <el-option v-for="(item, key) in PicTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="提示词ID" prop="promtId">
          <el-input v-model.number="searchInfo.promtId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="cookie类型" prop="cookieType">
          <el-select v-model="searchInfo.cookieType" clearable placeholder="请选择"
            @clear="() => { searchInfo.cookieType = undefined }">
            <el-option v-for="(item, key) in CookieTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="管理用户ID" prop="systemUserId">
          <el-input v-model.number="searchInfo.systemUserId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="项目状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择"
            @clear="() => { searchInfo.status = undefined }">
            <el-option v-for="(item, key) in SystemStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
        <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="Project_SystemProject" />
        <ExportExcel v-auth="btnAuth.exportExcel" template-id="Project_SystemProject" />
        <ImportExcel v-auth="btnAuth.importExcel" template-id="Project_SystemProject" @on-success="getTableData" />
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange" @sort-change="sortChange">
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column label="标题列表" prop="titleList" width="200">
          <template #default="scope">
            <ArrayCtrl v-model="scope.row.titleList" />
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="图片类型" prop="picType" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.picType, PicTypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="提示词ID" prop="promtId" width="120" />
        <el-table-column sortable align="left" label="cookie类型" prop="cookieType" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.cookieType, CookieTypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="管理用户ID" prop="systemUserId" width="120" />
        <el-table-column sortable align="left" label="项目状态" prop="status" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.status, SystemStatusOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button"
              @click="getDetails(scope.row)"><el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button"
              @click="updateSystemProjectFunc(scope.row)">变更</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete"
              @click="deleteRow(scope.row)">删除</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="flag" class="table-button"
              @click="SyncTitleFunc(scope.row)">同步标题</el-button>
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
        <el-form-item label="标题列表:" prop="titleList">
          <ArrayCtrl v-model="formData.titleList" editable />
        </el-form-item>
        <el-form-item label="图片类型:" prop="picType">
          <el-select v-model="formData.picType" placeholder="请选择图片类型" style="width:100%" :clearable="true">
            <el-option v-for="(item, key) in PicTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="提示词ID:" prop="promtId">
          <el-input v-model.number="formData.promtId" :clearable="true" placeholder="请输入提示词ID" />
        </el-form-item>
        <el-form-item label="cookie类型:" prop="cookieType">
          <el-select v-model="formData.cookieType" placeholder="请选择cookie类型" style="width:100%" :clearable="true">
            <el-option v-for="(item, key) in CookieTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="管理用户ID:" prop="systemUserId">
          <el-input v-model.number="formData.systemUserId" :clearable="true" placeholder="请输入管理用户ID" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="标题列表">
          <ArrayCtrl v-model="detailFrom.titleList" />
        </el-descriptions-item>
        <el-descriptions-item label="图片类型">
          {{ detailFrom.picType }}
        </el-descriptions-item>
        <el-descriptions-item label="提示词ID">
          {{ detailFrom.promtId }}
        </el-descriptions-item>
        <el-descriptions-item label="cookie类型">
          {{ detailFrom.cookieType }}
        </el-descriptions-item>
        <el-descriptions-item label="管理用户ID">
          {{ detailFrom.systemUserId }}
        </el-descriptions-item>
        <el-descriptions-item label="项目当前状态">
          {{ detailFrom.status }}
        </el-descriptions-item>
        
      </el-descriptions>
    </el-drawer>

    <el-drawer destroy-on-close size="500" v-model="syncTitleDialogVisible" :show-close="true"
      :before-close="closesyncTitleDialog">
      <template #header>
        <span class="text-lg">同步标题</span>
      </template>
      <el-form :model="formData" label-position="top" ref="syncTitleFormRef" label-width="80px">
        <el-form-item label="标题列表:" prop="titleList">
          <ArrayCtrl v-model="formData.titleList" editable />
        </el-form-item>
        <el-form-item label="上传Excel:" prop="titleList">
          <el-upload class="upload" 
            :action="`${getBaseUrl()}/fileUploadAndDownload/upload?noSave=1`"
            :on-change="onChange" 
            :show-file-list="true"
            >
            <el-button type="primary">点击上传Excel文件</el-button>
            <template #tip>
              <div class="el-upload__tip">
                仅支持上传 .xlsx 文件
              </div>
            </template>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button @click="SyncSyncTitleDialog">同 步</el-button>
          <el-button type="primary" @click="enterSyncTitleDialog">确 定</el-button>
          <el-button @click="closesyncTitleDialog">取 消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createSystemProject,
  deleteSystemProject,
  deleteSystemProjectByIds,
  updateSystemProject,
  findSystemProject,
  SyncTitle,
  getSystemProjectList
} from '@/api/Project/sysProject'
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'
// Excel解析组件
import * as XLSX from 'xlsx';

// 全量引入格式化工具，请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { getBaseUrl } from '@/utils/format'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'

defineOptions({
  name: 'SystemProject'
})
// 按钮权限实例化
const btnAuth = useBtnAuth()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const PicTypeOptions = ref([])
const CookieTypeOptions = ref([])
const SystemStatusOptions = ref([]);

const formData = ref({
  titleList: [],
  picType: '',
  promtId: undefined,
  cookieType: '',
  systemUserId: undefined,
  status: '',
})

// 验证规则
const rule = reactive({
  picType: [{
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
  promtId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  cookieType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }],
  systemUserId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  status: [{
    required: false,
    message: '状态为必填项',
    trigger: ['input', 'blur'],
  }],
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
const syncTitleFormRef = ref();
const tempTitleList = ref([]); // 临时变量用于存储 Excel 中的标题列表

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    titleList: 'title_list',
    picType: 'pic_type',
    promtId: 'promt_id',
    cookieType: 'cookie_type',
    systemUserId: 'system_user_id',
    status:'status',
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
  const table = await getSystemProjectList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  PicTypeOptions.value = await getDictFunc('PicType')
  CookieTypeOptions.value = await getDictFunc('CookieType')
  SystemStatusOptions.value = await getDictFunc('SystemStatus')

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
    deleteSystemProjectFunc(row)
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
    const res = await deleteSystemProjectByIds({ IDs })
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
const updateSystemProjectFunc = async (row) => {
  const res = await findSystemProject({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteSystemProjectFunc = async (row) => {
  const res = await deleteSystemProject({ ID: row.ID })
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

// 同步标题弹窗控制标记
const syncTitleDialogVisible = ref(false)

// 打开同步标题弹窗显示数据
const SyncTitleFunc = async (row) => {
  const res = await findSystemProject({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data
    syncTitleDialogVisible.value = true
  }
}

const closesyncTitleDialog = () => {
  syncTitleDialogVisible.value = false
}


// 文件选择回调
const onChange = (file, _) => {
  ReadExcel(file)
}



// 读取 Excel 文件
const ReadExcel = (file) => {
  try {
    const reader = new FileReader();
    reader.onload = (event) => {
      const data = new Uint8Array(event.target.result);
      const workbook = XLSX.read(data, { type: 'array' });
      // 这里可以进一步处理工作簿，例如获取工作表等
      const worksheet = workbook.Sheets[workbook.SheetNames[0]];
      const titleList = XLSX.utils.sheet_to_json(worksheet, { header: 1 });
      // 处理文件中的标题列表
      tempTitleList.value = titleList; // 将读取到的标题列表存储到tempTitleList
    };
    reader.readAsArrayBuffer(file.raw); // 使用readAsArrayBuffer
  } catch (error) {
    console.error('Error in ReadExcel:', error);
    ElMessage({ type: 'error', message: error.message });
  }
};




// 同步弹窗同步
const SyncSyncTitleDialog = async () => {

  try {
    const valid = await syncTitleFormRef.value?.validate(); // 直接获取表单验证结果
    if (!valid) throw new Error('表单验证未通过');

    // 将读取到的标题列表赋值给 formData.titleList
    formData.value.titleList = tempTitleList.value;

    // 进行同步操作
    const res = await SyncTitle(formData.value);
    if (res.code !== 0) {
      throw new Error(res.msg || '同步失败');
    }

    ElMessage({ type: 'success', message: '同步成功' });

  } catch (error) {
    console.error('Error in SyncSyncTitleDialog:', error); // 捕获错误
    ElMessage({ type: 'error', message: error.message });
  } finally {
    getTableData(); // 确保无论成功与否都调用获取数据
    formData.value = {
    titleList: [],
    picType: '',
    promtId: undefined,
    cookieType: '',
    systemUserId: undefined,
  }
  }
};

// 同步弹窗确定
const enterSyncTitleDialog = async () => {

  // try {
  //   const valid = await syncTitleFormRef.value?.validate(); // 直接获取表单验证结果
  //   if (!valid) throw new Error('表单验证未通过');

  //   const res = await SyncTitle(formData.value);
  //   if (res.code !== 0) {
  //     throw new Error(res.msg || '写入成功');
  //   }

  //   ElMessage({ type: 'success', message: '写入成功' });
    closesyncTitleDialog();

  // } catch (error) {
  //   console.error('Error in enterSyncTitleDialog:', error); // 捕获错误
  //   ElMessage({ type: 'error', message: error.message });
  // } finally {
  //   getTableData();
  //   formData.value = {
  //   titleList: [],
  //   picType: '',
  //   promtId: undefined,
  //   cookieType: '',
  //   systemUserId: undefined,
  //   status: '',
  // }
  // }
  
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
    titleList: [],
    picType: '',
    promtId: undefined,
    cookieType: '',
    systemUserId: undefined,
    status: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
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
      closeDialog()
      getTableData()
    }
  })
}

const detailFrom = ref({})
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async (row) => {
  const res = await findSystemProject({ ID: row.ID })
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

<style></style>
