import { useEffect, useState, useCallback } from 'react';
import { message } from 'antd';
import { getSurvey, listQuestion, getNext } from '@/services/ant-design-pro/survey';

/**
 * 地理定位 hook — 获取用户经纬度
 */
export function useGeolocation() {
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
          message.error('获取地理位置失败，请检查您的设置');
        },
      );
    } else {
      message.error('您的浏览器不支持地理位置功能');
    }
  }, []);

  return { latitude, longitude };
}

/**
 * 问卷数据加载 hook
 */
export function useSurveyLoader(surveyId: number) {
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
      message.error(error.message || '加载问卷数据失败');
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
export function useRespondentSN(sn: string | null) {
  const [generateRandom, setGenerateRandom] = useState('');

  useEffect(() => {
    if (sn) {
      setGenerateRandom(sn);
      getNext({ sn }).then((res) => {
        // 服务端返回下一个答题序号，组件自行处理跳转
        return res;
      }).catch(() => {
        message.error('恢复答题进度失败');
      });
    } else {
      setGenerateRandom(
        Math.random().toString(36).substring(2, 20),
      );
    }
  }, [sn]);

  return { generateRandom, setGenerateRandom };
}
