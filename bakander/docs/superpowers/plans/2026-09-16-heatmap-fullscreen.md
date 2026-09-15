# 地图热力图全屏 + 自适应范围 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 统计页地图弹窗改为全屏，热力图自动适配数据范围，支持一键回到全貌，热力标尺跟随数据。

**Architecture:** 纯逻辑（坐标解析、热力最大值）抽到独立 `heatmap-utils.ts` 并用 jest 做 TDD；`heatmap.tsx` 只做天地图 API 的接线（setViewport 适配视野、动态 max、回到全貌按钮、resize 适配）；`index.tsx` 只改 Modal 全屏样式。不新增依赖，不改后端。

**Tech Stack:** React 18 + TypeScript、UmiJS Max（Ant Design Pro 6）、antd 5.27 Modal、天地图 JS API v4.0、jest（@umijs/max/test）

**Spec:** `bakander/docs/superpowers/specs/2026-09-16-heatmap-fullscreen-design.md`

**重要背景：**
- 前端代码在 `survey-pro/` 目录；git 仓库根在 `D:\Project\kcsrsing\kcers-survey`（bakander 与 survey-pro 是同一个仓库的两个子目录），所有 git 命令在该仓库根执行，提交路径带 `survey-pro/` 前缀
- 提交只 add 本计划涉及的文件（仓库里有无关的已修改文件，如 `.idea/workspace.xml`、`bakander/config/config.yaml`，不要碰）
- 地图组件依赖天地图外部脚本（`window.T` 全局对象、canvas），无法在 jsdom 里做组件级单测（jsdom 无 canvas，alert 会抛错），因此单测只覆盖纯逻辑 utils，组件行为用 `npm run tsc` + 浏览器人工验证

## 文件结构

| 文件 | 操作 | 职责 |
|------|------|------|
| `survey-pro/src/pages/survey/statistics/components/heatmap-utils.ts` | 创建 | 纯函数：`parsePoints`（解析/过滤合法经纬度）、`computeHeatMax`（热力最大值） |
| `survey-pro/src/pages/survey/statistics/components/heatmap-utils.test.ts` | 创建 | 上述两个函数的 jest 单测 |
| `survey-pro/src/pages/survey/statistics/components/heatmap.tsx` | 修改 | 地图接线：setViewport 自适应、动态 max、回到全貌按钮、checkResize、容器 100% |
| `survey-pro/src/pages/survey/statistics/index.tsx` | 修改 | Modal 全屏样式 |

---

### Task 1: 热力图纯函数工具 + 单测（TDD）

**Files:**
- Create: `survey-pro/src/pages/survey/statistics/components/heatmap-utils.ts`
- Test: `survey-pro/src/pages/survey/statistics/components/heatmap-utils.test.ts`

- [ ] **Step 1: 先写失败的单测**

创建 `survey-pro/src/pages/survey/statistics/components/heatmap-utils.test.ts`：

```ts
import { computeHeatMax, parsePoints } from './heatmap-utils';

describe('parsePoints', () => {
  it('returns empty array for null/undefined/empty data', () => {
    expect(parsePoints(null as any)).toEqual([]);
    expect(parsePoints(undefined)).toEqual([]);
    expect(parsePoints([])).toEqual([]);
  });

  it('parses lng/lat strings to numbers', () => {
    const data = [
      { lng: '108.5', lat: '34.2' },
      { lng: '110.1', lat: '35.8' },
    ];
    expect(parsePoints(data)).toEqual([
      { lng: 108.5, lat: 34.2 },
      { lng: 110.1, lat: 35.8 },
    ]);
  });

  it('filters out non-numeric values', () => {
    const data = [
      { lng: 'abc', lat: '34.2' },
      { lng: '108.5', lat: '' },
      { lng: '109.0', lat: '33.9' },
    ];
    expect(parsePoints(data)).toEqual([{ lng: 109.0, lat: 33.9 }]);
  });
});

describe('computeHeatMax', () => {
  it('returns 1 for empty data', () => {
    expect(computeHeatMax([])).toBe(1);
    expect(computeHeatMax(null as any)).toBe(1);
  });

  it('returns the max count', () => {
    expect(computeHeatMax([{ count: 3 }, { count: 10 }, { count: 7 }])).toBe(10);
  });

  it('returns 1 when all counts are zero', () => {
    expect(computeHeatMax([{ count: 0 }])).toBe(1);
  });
});
```

- [ ] **Step 2: 运行测试确认失败**

在 `D:\Project\kcsrsing\kcers-survey\survey-pro` 下运行：

```bash
npx jest src/pages/survey/statistics/components/heatmap-utils.test.ts
```

Expected: FAIL，报错 `Cannot find module './heatmap-utils'`

- [ ] **Step 3: 实现工具函数**

创建 `survey-pro/src/pages/survey/statistics/components/heatmap-utils.ts`：

```ts
export interface HeatPoint {
  lng: string;
  lat: string;
}

export function parsePoints(
  data: HeatPoint[] | null | undefined,
): Array<{ lng: number; lat: number }> {
  if (!data) return [];
  return data
    .map((d) => ({ lng: parseFloat(d.lng), lat: parseFloat(d.lat) }))
    .filter((p) => Number.isFinite(p.lng) && Number.isFinite(p.lat));
}

export function computeHeatMax(data: Array<{ count: number }> | null | undefined): number {
  if (!data || data.length === 0) return 1;
  const max = data.reduce((m, d) => Math.max(m, d.count || 0), 0);
  return max > 0 ? max : 1;
}
```

- [ ] **Step 4: 运行测试确认通过**

```bash
npx jest src/pages/survey/statistics/components/heatmap-utils.test.ts
```

Expected: PASS，8 个用例全部通过

- [ ] **Step 5: 提交**

在仓库根 `D:\Project\kcsrsing\kcers-survey` 下：

```bash
git add survey-pro/src/pages/survey/statistics/components/heatmap-utils.ts survey-pro/src/pages/survey/statistics/components/heatmap-utils.test.ts
git commit -m "feat(survey-pro): 地图热力图工具函数（坐标解析/热力最大值）及单测"
```

---

### Task 2: heatmap.tsx 地图接线（自适应视野/动态标尺/回到全貌/窗口适配）

**Files:**
- Modify: `survey-pro/src/pages/survey/statistics/components/heatmap.tsx`（整文件重写，83 行）

- [ ] **Step 1: 重写组件**

用以下完整内容覆盖 `survey-pro/src/pages/survey/statistics/components/heatmap.tsx`：

```tsx
import { Button } from 'antd';
import React, { useEffect, useRef } from 'react';
import { computeHeatMax, parsePoints } from './heatmap-utils';

function isSupportCanvas() {
  const elem = document.createElement('canvas');
  return !!(elem.getContext && elem.getContext('2d'));
}

export const HeatMap = (props: { data: any[] }) => {
  const { data } = props;
  const mapRef = useRef<HTMLDivElement>(null);
  const mapInstance = useRef<any>(null);
  const heatmapInstance = useRef<any>(null);
  const scriptsLoaded = useRef(false);

  if (!isSupportCanvas()) {
    alert('热力图目前只支持有canvas支持的浏览器,您所使用的浏览器不能使用热力图功能~');
  }

  const fitToData = () => {
    const map = mapInstance.current;
    const T = (window as any).T;
    if (!map || !T) return;
    const points = parsePoints(data);
    if (points.length === 0) {
      map.centerAndZoom(new T.LngLat(108.95, 34.27), 4);
      return;
    }
    if (points.length === 1) {
      map.centerAndZoom(new T.LngLat(points[0].lng, points[0].lat), 12);
      return;
    }
    map.setViewport(points.map((p) => new T.LngLat(p.lng, p.lat)));
  };

  useEffect(() => {
    const loadScript = (src: string) => new Promise<void>((resolve, reject) => {
      if (document.querySelector(`script[src="${src}"]`)) {
        resolve();
        return;
      }
      const script = document.createElement('script');
      script.type = 'text/javascript';
      script.src = src;
      script.addEventListener('load', () => resolve());
      script.addEventListener('error', () => reject(new Error(`Failed to load ${src}`)));
      document.body.appendChild(script);
    });

    const createMap = () => {
      if (!mapRef.current) return;
      const T = (window as any).T;

      const map = new T.Map(mapRef.current);
      mapInstance.current = map;
      map.centerAndZoom(new T.LngLat(108.95, 34.27), 4);

      const heatmapOverlay = new T.HeatmapOverlay({
        radius: 30,
        maxOpacity: 0.9,
        minOpacity: 0.2,
        blur: 0.75,
        gradient: {
          0.0: 'blue',
          0.3: 'cyan',
          0.5: 'lime',
          0.7: 'yellow',
          0.85: 'orange',
          1.0: 'red',
        },
      });
      map.addOverLay(heatmapOverlay);
      heatmapInstance.current = heatmapOverlay;

      if (data && data.length > 0) {
        heatmapOverlay.setDataSet({ data, max: computeHeatMax(data) });
        heatmapOverlay.show();
      }
      fitToData();
    };

    if (scriptsLoaded.current) {
      // Scripts already loaded, just update data on existing map
      if (heatmapInstance.current && data && data.length > 0) {
        heatmapInstance.current.setDataSet({ data, max: computeHeatMax(data) });
        heatmapInstance.current.show();
      }
      return;
    }

    scriptsLoaded.current = true;

    loadScript('https://api.tianditu.gov.cn/api?v=4.0&tk=516e46ec670dc4149ad67ed5020d99fd')
      .then(() => loadScript('/scripts/HeatmapOverlay.js'))
      .then(createMap)
      .catch(console.error);

    const onResize = () => {
      mapInstance.current?.checkResize();
    };
    window.addEventListener('resize', onResize);
    return () => window.removeEventListener('resize', onResize);
  }, [data]);

  return (
    <div style={{ position: 'relative', width: '100%', height: '100%' }}>
      <div ref={mapRef} id="mapDiv" style={{ width: '100%', height: '100%' }} />
      <Button
        size="small"
        onClick={fitToData}
        style={{ position: 'absolute', top: 8, right: 8, zIndex: 10 }}
      >
        回到全貌
      </Button>
    </div>
  );
};
```

要点说明（实施时不要偏离）：
- 天地图 v4.0 `map.setViewport(点数组)` 自动调整中心+缩放包住所有点；`map.checkResize()` 在容器尺寸变化后重算
- 单点退化 `centerAndZoom(点, 12)`；无合法点保持原默认视野 (108.95, 34.27) zoom 4
- 热力 `max` 从写死 300 改为 `computeHeatMax(data)`
- resize 监听只在首次建图路径注册一次（early-return 分支不重复注册），组件卸载时移除
- `fitToData` 用最新 `data` 闭包，按钮点击与首次建图共用同一逻辑

- [ ] **Step 2: 类型检查**

在 `survey-pro` 下运行：

```bash
npm run tsc
```

Expected: 无 heatmap 相关报错（若输出与本改动无关的历史错误，忽略并在汇报中说明）

- [ ] **Step 3: 提交**

```bash
git add survey-pro/src/pages/survey/statistics/components/heatmap.tsx
git commit -m "feat(survey-pro): 热力图自适应数据范围、动态标尺、回到全貌按钮与窗口适配"
```

---

### Task 3: 统计页弹窗全屏

**Files:**
- Modify: `survey-pro/src/pages/survey/statistics/index.tsx:187-203`（Modal 部分）

- [ ] **Step 1: 修改 Modal 属性**

把 `survey-pro/src/pages/survey/statistics/index.tsx` 中的 Modal 从：

```tsx
      <Modal
        title={<p>地图</p>}
        centered
        width={{
          xs: '90%',
          sm: '80%',
          md: '70%',
          lg: '70%',
          xl: '70%',
          xxl: '80%',
        }}
        footer={<></>}
        open={openheatmap}
        onCancel={() => setOpenheatmap(false)}
      >
        {openheatmapdata ? <HeatMap data={openheatmapdata} /> : null}
      </Modal>
```

改为：

```tsx
      <Modal
        title={<p>地图</p>}
        width="100vw"
        style={{ top: 0, maxWidth: '100vw', paddingBottom: 0 }}
        styles={{ body: { height: 'calc(100vh - 120px)', padding: 0 } }}
        footer={<></>}
        open={openheatmap}
        onCancel={() => setOpenheatmap(false)}
      >
        {openheatmapdata ? <HeatMap data={openheatmapdata} /> : null}
      </Modal>
```

要点：`styles.body` 是 antd 5.10+ 的语义化 API（项目 antd ^5.27.3 支持）；高度扣 120px 给标题栏和边距；HeatMap 容器已是 100% 高，自动撑满。

- [ ] **Step 2: 类型检查**

```bash
npm run tsc
```

Expected: 无本次改动相关报错

- [ ] **Step 3: 提交**

```bash
git add survey-pro/src/pages/survey/statistics/index.tsx
git commit -m "feat(survey-pro): 统计页地图弹窗改为全屏"
```

---

### Task 4: 浏览器人工验证

**Files:** 无代码改动

- [ ] **Step 1: 启动前端**

在 `survey-pro` 下运行（保持前台运行）：

```bash
npm run dev
```

Expected: umi 编译完成后提示本地地址（默认 http://localhost:8000）

- [ ] **Step 2: 验证主流程**

浏览器打开 `http://localhost:8000/survey/3/statistics`，点「地图」，逐项确认：

1. 弹窗占满浏览器窗口，地图无留白/无滚动条
2. 初始视野自动包住全部调研点（能看到所有热力斑块，即"全貌"）
3. 热力颜色有层次（不同密集度颜色可区分，不再一片同色）
4. 鼠标滚轮放大后能看清局部细节
5. 点右上角「回到全貌」按钮，视野恢复初始全貌
6. 缩放浏览器窗口大小，地图正常重绘不错位
7. 关闭弹窗再重新打开，地图正常
8. 控制台无报错（Network 里天地图 API 与 /scripts/HeatmapOverlay.js 加载成功）

- [ ] **Step 3: 验证边界情况**

打开另一个数据极少或无数据的问卷统计页（如 `http://localhost:8000/survey/1/statistics` 或其他 id），点「地图」：

1. 无数据：显示默认视野（全国），不报错，无热力
2. 仅 1 个点（如存在此类问卷）：初始视野定位到该点附近（zoom 12），不无限放大

- [ ] **Step 4: 完成**

全部通过后在会话中汇报验证结果（截图或逐项结论）。本任务无提交。

---

## Self-Review 记录

- Spec 覆盖：设计文档 5 个条目 → Task 1（动态标尺的纯逻辑）、Task 2（自适应/按钮/checkResize/动态标尺接线）、Task 3（全屏弹窗）、Task 4（含边界情况的浏览器验证）；spec 中"仅前端、不动后端"在文件结构中体现
- 占位符扫描：无 TBD/TODO，所有代码步骤含完整代码
- 类型一致性：`parsePoints`/`computeHeatMax` 在 Task 1 定义、Task 2 引用的签名一致；`fitToData` 在组件内定义并被按钮与建图流程复用，命名一致