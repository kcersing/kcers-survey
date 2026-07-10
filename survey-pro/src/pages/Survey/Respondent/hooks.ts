import { useEffect, useState, useCallback } from 'react';
import { message as staticMessage } from 'antd';
import { getSurvey, listQuestion, getNext } from '@/services/ant-design-pro/survey';

type MessageApi = typeof staticMessage;

/**
 * 地理定位 hook — 获取用户经纬度
 */
export function useGeolocation(msgApi: MessageApi = staticMessage) {
  const [latitude, setLatitude] = useState<number | null>(null);
  const [longitude, setLongitude] = useState<number | null>(null);

  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          setLatitude(position.coords.latitude);
          setLongitude(position.coords.longitude);
        },
        () => {
          msgApi.error('获取地理位置失败，请检查您的设置');
        },
      );
    } else {
      msgApi.error('您的浏览器不支持地理位置功能');
    }
  }, []);

  return { latitude, longitude };
}

/**
 * 问卷数据加载 hook
 */
export function useSurveyLoader(surveyId: number, msgApi: MessageApi = staticMessage) {
  const [survey, setSurvey] = useState<API.Survey>({});
  const [questions, setQuestions] = useState<API.Questions[]>([]);

  const load = useCallback(async () => {
    try {
      if (!surveyId) return;
      const [surveyData, questionsData] = await Promise.all([
        getSurvey({ id: surveyId }),
        listQuestion({ surveyId }),
      ]);
      setSurvey(surveyData.data);
      setQuestions(questionsData.data);
    } catch (error: any) {
      msgApi.error(error.message || '加载问卷数据失败');
    }
  }, [surveyId]);

  useEffect(() => {
    load();
  }, [load]);

  return { survey, questions, reload: load };
}

/**
 * 答题断点续答 SN hook
 */
export function useRespondentSN(sn: string | null, msgApi: MessageApi = staticMessage) {
  const [generateRandom, setGenerateRandom] = useState('');

  useEffect(() => {
    if (sn) {
      setGenerateRandom(sn);
      getNext({ sn }).then((res) => {
        return res;
      }).catch(() => {
        msgApi.error('恢复答题进度失败');
      });
    } else {
      setGenerateRandom(
        Math.random().toString(36).substring(2, 20),
      );
    }
  }, [sn]);

  return { generateRandom, setGenerateRandom };
}

/**
 * 从扁平列表构建树形结构（基于 parentId），children 会被正确填充。
 */
export function buildTree(flatList: API.Questions[]): API.Questions[] {
  if (!flatList?.length) return [];
  const idMap = new Map<number, API.Questions>();
  const roots: API.Questions[] = [];

  for (const q of flatList) {
    idMap.set(q.id, { ...q, children: q.children?.length ? [...q.children] : [] });
  }
  for (const q of flatList) {
    const node = idMap.get(q.id)!;
    if (q.parentId && idMap.has(q.parentId)) {
      const parent = idMap.get(q.parentId)!;
      if (!parent.children) parent.children = [];
      parent.children.push(node);
    } else {
      roots.push(node);
    }
  }
  return roots;
}

function flatQuestions(tree: API.Questions[]): API.Questions[] {
  if (!tree?.length) return [];
  const result: API.Questions[] = [];
  for (const q of tree) {
    if (!q) continue;
    if (q.type === 'h2' || q.type === 'h3') {
      if (q.children?.length > 0) {
        result.push(...flatQuestions(q.children));
      }
    } else {
      result.push(q);
    }
  }
  return result;
}

export interface StepItem {
  /** 步骤唯一标识 */
  name: string;
  /** 步骤上要渲染的所有题目 */
  questions: API.Questions[];
  /** 步骤顶部的 h2 标题 */
  h2heading?: API.Questions;
  /** 步骤顶部的 h3 标题 */
  h3heading?: API.Questions;
}

/**
 * 从树形题目构建步骤列表。
 * - h2：子题各成一页，顶部显示 h2 标题
 * - h3：所有子题合在一页，顶部显示 h3 标题
 * - 普通题目：各自一页
 */
export function buildSteps(tree: API.Questions[]): StepItem[] {
  if (!tree?.length) return [];
  const result: StepItem[] = [];

  for (const q of tree) {
    if (!q) continue;

    if (q.type === 'h2') {
      // h2: 子题各自成步骤，带 h2 标题
      const children = buildSteps(q.children || []);
      for (const child of children) {
        child.h2heading = child.h2heading || q;
        result.push(child);
      }
    } else if (q.type === 'h3') {
      // h3: 所有子题合在一个步骤
      const flatChildren = flatQuestions(q.children || []);
      result.push({
        name: `h3-${q.id}`,
        questions: flatChildren,
        h3heading: q,
      });
    } else {
      // 普通题目
      result.push({
        name: String(q.id),
        questions: [q],
      });
    }
  }

  return result;
}

/**
 * 在树形题目中按 ID 查找题目
 */
export function findInTree(tree: API.Questions[], id: number): API.Questions | null {
  for (const q of tree) {
    if (q.id === id) return q;
    if (q.children?.length > 0) {
      const found = findInTree(q.children, id);
      if (found) return found;
    }
  }
  return null;
}
