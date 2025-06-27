import React, { useState, useReducer, useEffect, useRef, memo, useCallback } from 'react';
import {
  Button,
  Typography,
  message,
  Spin,
} from 'antd';
import { ArrowLeftOutlined, } from '@ant-design/icons';
import { useNavigate, useParams } from 'umi';
import {
  ProForm,
  ProCard ,
  ProFormDigit,
  ProFormText,
  ProFormUploadDragger,
  ProFormUploadButton,
  ProFormRate,
  ProFormRadio,
  ProFormCheckbox,
  ProFormTextArea,
  ProFormDatePicker,
} from '@ant-design/pro-components';
import './survey.css';

import { getSurvey, listQuestion } from '@/services/ant-design-pro/survey';
const { Title, Text } = Typography;


// 问题渲染组件 - 使用 memo 优化
const QuestionRenderer = memo(({
                                 question,
                                 depth,
                                 form,
                                 answers,
                                 onUploadStart,
                                 onUploadSuccess,
                                 onUploadError
                               }) => {

console.log(question)
  // 普通问题
  return (
    <div className={`question question-depth-${depth}`}>
      <ProForm.Item
        name={`question${question.id}`}
        label={question.content}
        rules={question.required ? [{ required: true, message: '此字段为必填项' }] : []}
        initialValue={answers[question.id]?.[`question${question.id}`]}
      >
        {renderQuestionControl(question, depth, {
          onUploadStart,
          onUploadSuccess,
          onUploadError
        })}
      </ProForm.Item>
    </div>
  );
}, (prevProps, nextProps) => {
  // 浅比较问题和答案，避免不必要的渲染
  return prevProps.question.id === nextProps.question.id &&
    prevProps.answers[prevProps.question.id] === nextProps.answers[nextProps.question.id];
});

// 问题控件渲染
const renderQuestionControl = (
  question:API.Questions,
  depth: number,
  uploadHandlers: {
    onUploadStart: (type: string) => void;
    onUploadSuccess: (type: string, file: any) => void;
    onUploadError: (type: string, error: any) => void;
  }
) => {
  const { onUploadStart, onUploadSuccess, onUploadError } = uploadHandlers;
  console.log(question)

  switch (question.type) {
    case 'single_choice':
      return (
      <ProFormRadio.Group
        name={question.id}
        options={question.options.map(option => ({
          value:option.content,
          label: option.content,
        }))}
      />
      );

    case 'multiple_choice':
      return (
      <ProFormCheckbox.Group
        name={question.id}
        options={question.options.map(option => ({
          value:option.content,
          label: option.content,
        }))}
      />
      );

    case 'text':
      return <ProFormTextArea name={question.id} label="名称"  placeholder="请输入内容..." />;

    case 'number':
      return <ProFormDigit placeholder="请输入数字"
                           name={question.id}
                           min={1}
                           max={10} />;

    case 'date':
      return <ProFormDatePicker  name={question.id} placeholder="请选择日期" />;

    case 'rate':
      return (
        <ProFormRate name={question.id} label="Rate" />
      );

    case 'uploadImage':
      return (
        <ProFormUploadButton  name={question.id} label="Upload" />
        // <ProFormUploadDragger
        //   name="file"
        //   className="upload-image"
        //   onStart={() => onUploadStart('image')}
        //   onSuccess={(info) => onUploadSuccess('image', info)}
        //   onError={(error) => onUploadError('image', error)}
        // >
        //   <Button icon={<PlusOutlined />}>上传图片</Button>
        // </ProFormUploadDragger>
      );

    case 'uploadFile':
      return (
        <ProFormUploadButton name={question.id} label="Upload" />
        // <ProFormUploadDragger
        //   name="file"
        //   className="upload-file"
        //   onStart={() => onUploadStart('file')}
        //   onSuccess={(info) => onUploadSuccess('file', info)}
        //   onError={(error) => onUploadError('file', error)}
        // >
        //   <Button icon={<PlusOutlined />}>上传文件</Button>
        // </ProFormUploadDragger>
      );
    case 'h2':
      return (
        <h3>{question.content}</h3>
      );
    case 'h3':
      return (
          <h4>{question.content}</h4>
      );

    default:
      return null;
  }
};

// 导航按钮组件 - 使用 memo 优化
const NavigationButtons = memo(({
                                  isFirstQuestion,
                                  isSubmitting,
                                  uploading,
                                  moveToPreviousQuestion,
                                  moveToNextQuestion,
                                  form
                                }) => {
  return (
    <div className="survey-buttons">
      <Button
        type="default"
        icon={<ArrowLeftOutlined />}
        disabled={isFirstQuestion || isSubmitting}
        onClick={moveToPreviousQuestion}
      >
        上一题
      </Button>
      <Button
        type="primary"
        onClick={moveToNextQuestion} // 直接调用导航函数，不依赖表单提交
        loading={isSubmitting || uploading}
        style={{ marginLeft: 16 }}
      >
        {isSubmitting ? '提交中...' : '下一步'}
      </Button>
    </div>
  );
});

// 调查状态类型
type SurveyState = {
  loading: boolean;
  showThankYou: boolean;
  isSubmitting: boolean;
  uploading: boolean;
};

// 调查动作类型
type SurveyAction =
  | { type: 'LOADING'; payload: boolean }
  | { type: 'SUBMITTING'; payload: boolean }
  | { type: 'UPLOADING'; payload: boolean }
  | { type: 'THANK_YOU'; payload: boolean };

// 调查状态 reducer
const surveyReducer = (state: SurveyState, action: SurveyAction): SurveyState => {
  switch (action.type) {
    case 'LOADING':
      return { ...state, loading: action.payload };
    case 'SUBMITTING':
      return { ...state, isSubmitting: action.payload };
    case 'UPLOADING':
      return { ...state, uploading: action.payload };
    case 'THANK_YOU':
      return { ...state, showThankYou: action.payload };
    default:
      return state;
  }
};

// 问卷调查组件
const Survey = () => {
  const [form] = ProForm.useForm();
  const [currentPath, setCurrentPath] = useState<number[]>([]);
  const [answers, setAnswers] = useState<Record<number, any>>({});

  // 使用 useReducer 管理复杂状态
  const [surveyState, dispatch] = useReducer(surveyReducer, {
    loading: true,
    showThankYou: false,
    isSubmitting: false,
    uploading: false,
  });


  const [questions, setQuestions] = useState<API.Questions[]>([]);
  const [survey, setSurvey] = useState<API.Survey>({});

  // // 预处理问题数据
  // const [processedQuestions, setProcessedQuestions] = useState({
  //   allQuestions: [] as API.Questions[],
  //   questionMap: new Map<number,API.Questions>(),
  //   pathMap: new Map<number, number[]>(),
  // });

  const navigate = useNavigate();
  const firstRender = useRef(true);
  const { id } = useParams();
  const surveyId = id ? parseInt(id) : 1;

  // 加载问卷和问题数据
  useEffect(() => {
    loadSurveyAndQuestions();
  }, []);


  // 设置初始路径 - 只考虑根问题
  useEffect(() => {
    if (questions.length > 0 && firstRender.current) {
      // 找到第一个非标题类型的根问题
      const firstQuestion = questions.find(q => q.type !== 'h2' && q.type !== 'h3');
      if (firstQuestion) {
        setCurrentPath([questions.indexOf(firstQuestion)]);
        firstRender.current = false;
      }
    }
  }, [questions]);

  // 加载问卷和问题数据
  const loadSurveyAndQuestions = async () => {
    try {
      if (!surveyId) return;

      dispatch({ type: 'LOADING', payload: true });

      const [surveyData, questionsData] = await Promise.all([
        getSurvey({ id: surveyId }),
        listQuestion({ surveyId: surveyId }),
      ]);

      // 处理数据格式
      const processedQuestions = (questionsData.data || []).map(question => ({
        ...question,
        required: question.required === 1 // 将数字转换为布尔值
      }));

      setSurvey(surveyData.data || {});
      setQuestions(processedQuestions);
    } catch (error: any) {
      console.error('加载问卷数据失败', error);
      message.error(error.message || '加载问卷数据失败');
    } finally {
      dispatch({ type: 'LOADING', payload: false });
    }
  };

  // 获取当前问题 - 只考虑根问题
  const getCurrentQuestion = useCallback((path: number[]) => {
console.log(path)
    if (!path || path.length === 0) {
      return questions[0];
    }

    return questions[path[0]] || null;

  }, [questions]);

  // 查找下一个根问题 - 简化版本
  const findNextRootQuestion = useCallback(() => {
    console.log(3)
    // 获取当前问题在根问题列表中的索引

    const currentIndex = currentPath.length > 0 ? currentPath[0] : -1;

    // 查找下一个非标题类型的根问题
    for (let i = currentIndex + 1; i < questions.length; i++) {
      console.log([i])
        return [i]; // 返回根问题的路径
    }

    return null; // 没有更多根问题
  }, [questions, currentPath]);

  // 查找上一个根问题 - 简化版本
  const findPrevRootQuestion = useCallback(() => {
    console.log(2)
    // 获取当前问题在根问题列表中的索引
    const currentIndex = currentPath.length > 0 ? currentPath[0] : 0;

    // 查找上一个非标题类型的根问题
    for (let i = currentIndex - 1; i >= 0; i--) {
      if (questions[i].type !== 'h2' && questions[i].type !== 'h3') {
        return [i]; // 返回根问题的路径
      }
    }

    return null; // 没有更多根问题
  }, [questions, currentPath]);

  // 处理跳题 - 简化版本，只考虑根问题
  const handleJumpRules = useCallback((question: API.Questions, values: any) => {
    if (!question.jumpRules || question.jumpRules.length === 0) return null;

    const selectedAnswer = values[`question${question.id}`];

    // 处理单选和多选的不同情况
    if (question.type === 'single_choice') {
      // 单选处理
      for (const rule of question.jumpRules) {
        if (rule.operators === 'equals' && String(selectedAnswer) === rule.answer) {
          // 找到目标问题在根问题中的索引
          const targetIndex = questions.findIndex(q => q.id === rule.nextQuestionId);
          if (targetIndex !== -1) {
            return [targetIndex];
          }
        }
      }
    } else if (question.type === 'multiple_choice' && Array.isArray(selectedAnswer)) {
      // 多选处理
      for (const rule of question.jumpRules) {
        // 检查是否包含特定选项
        if (rule.operators === 'includes' && selectedAnswer.includes(rule.answer)) {
          const targetIndex = questions.findIndex(q => q.id === rule.nextQuestionId);
          if (targetIndex !== -1) {
            return [targetIndex];
          }
        }
      }
    }

    return null;
  }, [questions]);

  // 移动到下一个问题 - 简化版本
  const moveToNextQuestion = useCallback(() => {
    console.log(currentPath)

    const currentQuestion = getCurrentQuestion(currentPath);
    if (!currentQuestion) {
      handleSurveySubmit();
      return;
    }

    // 处理跳题规则
    const jumpPath = handleJumpRules(currentQuestion, answers[currentQuestion.id] || {});
    if (jumpPath) {
      setCurrentPath(jumpPath);
      return;
    }

    // 查找下一个根问题
    const nextPath = findNextRootQuestion();
    if (nextPath) {
      setCurrentPath(nextPath);
    } else {
      // 如果没有更多问题，提交问卷
      handleSurveySubmit();
    }
  }, [currentPath, getCurrentQuestion, handleJumpRules, findNextRootQuestion, answers, handleSurveySubmit]);

  // 移动到上一个问题 - 简化版本
  const moveToPreviousQuestion = useCallback(() => {
    console.log(4)
    const prevPath = findPrevRootQuestion();
    if (prevPath) {
      setCurrentPath(prevPath);
    }
  }, [findPrevRootQuestion]);

  // 处理问题提交
  const handleQuestionSubmit = useCallback((values: any) => {
    console.log(5)
    const currentQuestion = getCurrentQuestion(currentPath);
    if (!currentQuestion) return;

    // 使用函数式更新确保状态一致性
    setAnswers(prev => ({
      ...prev,
      [currentQuestion.id]: values
    }));

    moveToNextQuestion();
  }, [currentPath, getCurrentQuestion, moveToNextQuestion]);

  // 处理问卷提交
  const handleSurveySubmit = useCallback(() => {
    console.log(6)
    dispatch({ type: 'SUBMITTING', payload: true });

    // 验证必填问题 - 只考虑根问题
    const requiredQuestions = questions.filter(
      q => q.type !== 'h2' && q.type !== 'h3' && q.required
    );

    const hasMissingAnswers = requiredQuestions.some(q => {
      const answer = answers[q.id]?.[`question${q.id}`];
      return answer === undefined || answer === null || (Array.isArray(answer) && answer.length === 0);
    });

    if (hasMissingAnswers) {
      message.error('请完成所有必填问题');
      dispatch({ type: 'SUBMITTING', payload: false });
      return;
    }

    // 实际项目中应替换为真实的提交逻辑
    setTimeout(() => {
      dispatch({ type: 'THANK_YOU', payload: true });
      dispatch({ type: 'SUBMITTING', payload: false });
    }, 1000);
  }, [answers, questions]);

  // 格式化进度显示 - 只考虑根问题
  const formatProgress = useCallback(() => {
    console.log(7)
    const allQuestions = questions.filter(q => q.type !== 'h2' && q.type !== 'h3');
    const answeredCount = allQuestions.filter(q =>
      answers[q.id] !== undefined &&
      answers[q.id][`question${q.id}`] !== undefined &&
      answers[q.id][`question${q.id}`] !== null
    ).length;

    return `${answeredCount}/${questions.length}`;
  }, [answers, questions]);






  // 处理文件上传
  const handleUploadStart = useCallback((type: string) => {
    dispatch({ type: 'UPLOADING', payload: true });
  }, []);

  const handleUploadSuccess = useCallback((type: string, file: any) => {
    dispatch({ type: 'UPLOADING', payload: false });
    message.success('上传成功');
  }, []);

  const handleUploadError = useCallback((type: string, error: any) => {
    dispatch({ type: 'UPLOADING', payload: false });
    message.error('上传失败: ' + error.message);
  }, []);

  // 导航调试日志
  useEffect(() => {
    console.log('Current Path:', currentPath);
    const currentQuestion = getCurrentQuestion(currentPath);
    if (currentQuestion) {
      console.log('Current Question:', currentQuestion.id, currentQuestion.content);
    }
  }, [currentPath, getCurrentQuestion]);

  // 渲染感谢页面
  const renderThankYou = () => {
    return (
      <div className="thank-you-container">
        <ProCard className="thank-you-card" bordered={false}>
          <p className="thank-you-icon" />
          <Title level={3}>感谢您参与调查！</Title>
          <Text>您的反馈对我们非常重要，我们将根据您的意见改进服务。</Text>
          <Button
            type="primary"
            onClick={() => navigate('/')}
            className="finish-button"
          >
            完成
          </Button>
        </ProCard>
      </div>
    );
  };

  if (surveyState.showThankYou) {
    return renderThankYou();
  }

  if (surveyState.loading || questions.length === 0) {
    return (
      <div className="loading-message">
        <Spin tip="加载问卷中..." />
      </div>
    );
  }

  console.log(currentPath)
  const currentQuestion = getCurrentQuestion(currentPath);


  const recursionQuestion =(currentQuestion,currentPath,form,answers,handleUploadStart,handleUploadSuccess,handleUploadError)=>{

    console.log(8)
    return (
      <>
      {currentQuestion && (
      <QuestionRenderer
        question={currentQuestion}
        depth={currentPath.length}
        form={form}
        answers={answers}
        onUploadStart={handleUploadStart}
        onUploadSuccess={handleUploadSuccess}
        onUploadError={handleUploadError}
      />
    )}

    {/* 渲染子问题 - 只在有子问题时显示 */}
    {currentQuestion && currentQuestion.children && currentQuestion.children.length > 0 && (
      <div className="children-questions">
        {currentQuestion.children.map((child, index) => (
          <div key={child.id} style={{ paddingLeft: '20px' }}>

            { recursionQuestion(child,currentPath,form,answers,handleUploadStart,handleUploadSuccess,handleUploadError)}

          </div>
        ))}
      </div>
    )} </>);
  }



  return (
    <div className="survey-container">
      <ProCard
        className="survey-card"
        title={currentQuestion?.content || survey.title || '问卷调查'}
        extra={<Text type="secondary">答题进度: {formatProgress()}</Text>}
        bordered={false}
      >
        <ProForm form={form} onFinish={handleQuestionSubmit} layout="vertical">

          {console.log(currentQuestion)}
          {currentQuestion && recursionQuestion(currentQuestion,currentPath,form,answers,handleUploadStart,handleUploadSuccess,handleUploadError)}

          {/*/!* 渲染子问题 - 只在有子问题时显示 *!/*/}
          {/*{currentQuestion && currentQuestion.children && currentQuestion.children.length > 0 && (*/}
          {/*  <div className="children-questions">*/}
          {/*    {currentQuestion.children.map((child, index) => (*/}
          {/*      <div key={child.id} style={{ paddingLeft: '20px' }}>*/}
          {/*        <QuestionRenderer*/}
          {/*          question={child}*/}
          {/*          depth={currentPath.length + 1}*/}
          {/*          form={form}*/}
          {/*          answers={answers}*/}
          {/*          onUploadStart={handleUploadStart}*/}
          {/*          onUploadSuccess={handleUploadSuccess}*/}
          {/*          onUploadError={handleUploadError}*/}
          {/*        />*/}
          {/*      </div>*/}
          {/*    ))}*/}
          {/*  </div>*/}
          {/*)}*/}

          <NavigationButtons
            isFirstQuestion={currentPath.length === 0}
            isSubmitting={surveyState.isSubmitting}
            uploading={surveyState.uploading}
            moveToPreviousQuestion={moveToPreviousQuestion}
            moveToNextQuestion={moveToNextQuestion}
            form={form}
          />
        </ProForm>
      </ProCard>
    </div>
  );
};

export default Survey;
