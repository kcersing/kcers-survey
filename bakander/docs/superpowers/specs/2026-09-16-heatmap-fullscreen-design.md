# 问卷统计地图热力图：全屏弹窗 + 自适应范围 设计文档

日期：2026-09-16
状态：已确认（等待用户审阅）

## 背景与问题

页面 `/survey/3/statistics` 中点击「地图」会打开一个宽度 70% 的弹窗，内部是天地图热力图（`survey-pro/src/pages/survey/statistics/components/heatmap.tsx`）。现状：

- 地图固定中心 (108.95, 34.27)、固定 zoom 4
- 热力最大值写死 `max: 300`
- 弹窗 70% 宽、地图容器写死 900px 高

调研点约几百个、省内/跨省分布。zoom 4 下热点糊成一团看不出分布；放大后又丢失全貌，且没有一键回到全貌的途径。

## 目标

1. 弹窗改为全屏（占满浏览器窗口）
2. 打开地图时自动适配视野，包住全部调研点（全貌）
3. 放大后可通过「回到全貌」按钮一键恢复
4. 热力颜色标尺跟随实际数据，全貌下也能看出分布差异

## 方案（仅前端，不动后端接口）

### 1. 弹窗改全屏 — `statistics/index.tsx`

Modal 调整为：

- `width="100vw"`，`style={{ top: 0, maxWidth: '100vw', paddingBottom: 0 }}`
- 内容区高度 `calc(100vh - 120px)`（120px 为标题栏 + 安全边距），padding 0
- 保留标题「地图」和右上角关闭按钮，footer 不变（`<></>`）

### 2. 地图自适应数据范围 — `heatmap.tsx`

创建地图后，用 `data` 计算边界：

- 解析每个点的 `parseFloat(lng/lat)`，取 min/max 边界
- 调用天地图 v4.0 `map.setViewport(点数组)`，自动定位到包住全部调研点的视野
- 边界情况：
  - 仅 1 个点：退化为 `centerAndZoom(该点, 12)`，避免过度放大
  - 0 个点：保持现状默认视野（中心 108.95, 34.27，zoom 4）

### 3. 「回到全貌」按钮 — `heatmap.tsx`

- 地图容器内右上角悬浮一个 antd Button（绝对定位、z-index 高于地图）
- 点击重新执行与初始化相同的 `setViewport`
- 边界数组存入 `useRef`，供按钮复用

### 4. 热力动态标尺 — `heatmap.tsx`

- `setDataSet({ data, max: 300 })` 中的 max 改为取数据实际最大值：`Math.max(...data.map(d => d.count))`（最小为 1）

### 5. 窗口尺寸变化适配 — `heatmap.tsx`

- 地图高度改为动态（100% 撑满弹窗），新增 window resize 监听，触发 `map.checkResize()`（天地图 v4.0 API），避免浏览器窗口缩放后地图渲染错位

## 涉及文件

| 文件 | 改动 |
|------|------|
| `survey-pro/src/pages/survey/statistics/index.tsx` | Modal 全屏样式 |
| `survey-pro/src/pages/survey/statistics/components/heatmap.tsx` | setViewport 自适应、动态 max、回到全貌按钮、checkResize |

后端 `GetSurveyResponseHeatmap` 接口与数据格式（lng/lat/count）不变，无需修改 thrift。

## 测试验证

- `npm run dev` 打开 `/survey/3/statistics`，点「地图」
- 验证：弹窗占满窗口；初始视野包住全部调研点；放大后点「回到全貌」可恢复；热力颜色有层次（不同 count 颜色可区分）；缩放浏览器窗口地图正常
- 边界：只有一个调研点的问卷、无数据问卷
- 类型检查：`npm run tsc`

## 非目标（YAGNI）

- 不做聚合点/气泡点方案（用户选择保留热力密度表达）
- 不做鹰眼缩略图
- 不调整热力 radius/gradient 参数（现状可接受）
- 不做独立地图页面