import React, { useState, useEffect, useRef } from 'react';
import {

  Form,
  Radio,
  Checkbox,
  Input,
  Button,
  Typography, message,

} from 'antd';
import { QuestionCircleOutlined,ArrowLeftOutlined } from '@ant-design/icons';
import { useNavigate } from 'umi';
import {
  ProCard,
  PageContainer,
  ProForm,
  ProFormDatePicker, ProFormDigit,
} from '@ant-design/pro-components';
import { useParams } from "react-router"
import {
  listQuestion,
  getSurvey,
} from '@/services/ant-design-pro/survey';


import './survey.css';

const { Title, Text } = Typography;

// 答题状态管理
const Survey = () => {

  const [form] = Form.useForm();
  const [currentSectionIndex, setCurrentSectionIndex] = useState(0);
  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [showThankYou, setShowThankYou] = useState(false);
  const navigate = useNavigate();
  const sectionRefs = useRef([]);
  // 保存答题进度
  const [answers, setAnswers] = useState({});

  const [loading, setLoading] = useState(false);
  const [questions, setQuestions] = useState<API.Questions[]>([]);
  const [survey, setSurvey] = useState<API.Survey>({});

  let params = useParams();
  const surveyId =parseInt(params.id)
  // 问卷数据


    const loadSurveyAndQuestions = async () => {
      try {
        // 并行加载问卷和问题
        const [surveyData, questionsData] = await Promise.all([
          getSurvey({id:surveyId}),
          listQuestion({surveyId:surveyId}),
        ]);
        setSurvey(surveyData.data);
        setQuestions(questionsData.data);
        setLoading(true);

      } catch (error) {
        message.error('加载问卷数据失败');
      } finally {
        setLoading(false);
      }
    };

    useEffect(() => {
      loadSurveyAndQuestions()
    }, []);

  // // 初始化ref数组
    useEffect(() => {
      sectionRefs.current = Array(questions.length).fill(null);
    }, [questions.length]);


  // 处理问题提交和跳题
  const handleQuestionSubmit = (values) => {
    // 保存答案
    const question = getCurrentQuestion();
    setAnswers(prev => ({
      ...prev,
      [question.id]: values
    }));

    // 检查跳题规则
    const jumpTo = checkJumpRules(question, values);
    if (jumpTo) {
      // 跳转到指定问题
      const { sectionIndex, questionIndex } = findQuestionIndexById(jumpTo);
      setCurrentSectionIndex(sectionIndex);
      setCurrentQuestionIndex(questionIndex);
      return;
    }

    // 移动到下一个问题
    moveToNextQuestion();
  };

  // 检查跳题规则
  const checkJumpRules = (question, values) => {
    const jumpRules = question.jumpRules;
    if (!jumpRules || jumpRules.questionId === 0) return null;

    // 简化处理：这里假设answer是选项的serial值
    const selectedAnswer = values[`question${question.id}`];
    if (jumpRules.operators === '包含' && jumpRules.answer.includes(selectedAnswer + '')) {
      return jumpRules.nextQuestionId;
    }
    return null;
  };

  // 根据ID查找问题索引
  const findQuestionIndexById = (questionId) => {
    for (let s = 0; s < questions.length; s++) {
      const section = questions[s];
      if (section.children) {
        for (let q = 0; q < section.children.length; q++) {
          if (section.children[q].id === questionId) {
            return { sectionIndex: s, questionIndex: q };
          }
        }
      }
    }
    return { sectionIndex: currentSectionIndex, questionIndex: currentQuestionIndex + 1 };
  };

  // 移动到下一个问题
  const moveToNextQuestion = () => {
    const section = questions[currentSectionIndex];
    if (!section.children) return;

    const nextQuestionIndex = currentQuestionIndex + 1;
    if (nextQuestionIndex < section.children.length) {
      setCurrentQuestionIndex(nextQuestionIndex);
    } else {
      // 本部分已完成，移动到下一部分
      const nextSectionIndex = currentSectionIndex + 1;
      if (nextSectionIndex < questions.length) {
        setCurrentSectionIndex(nextSectionIndex);
        setCurrentQuestionIndex(0);
      } else {
        // 所有问题已完成
        handleSurveySubmit();
      }
    }
  };
  // 移动到上一个问题
  const moveToPreviousQuestion = () => {
    const section = questions[currentSectionIndex];
    if (!section.children) return;

    const prevQuestionIndex = currentQuestionIndex - 1;
    if (prevQuestionIndex >= 0) {
      setCurrentQuestionIndex(prevQuestionIndex);
    } else {
      // 上一部分的最后一个问题
      const prevSectionIndex = currentSectionIndex - 1;
      if (prevSectionIndex >= 0) {
        const prevSection = questions[prevSectionIndex];
        setCurrentSectionIndex(prevSectionIndex);
        setCurrentQuestionIndex(prevSection.children.length - 1);
      }
    }
  };
  // 处理问卷提交
  const handleSurveySubmit = () => {
    setIsSubmitting(true);
    // 模拟提交数据
    setTimeout(() => {
      setShowThankYou(true);
      setIsSubmitting(false);
    }, 1500);
  };

  // 获取当前部分
  const getCurrentSection = () => {
    if (questions){
      return questions[currentSectionIndex];
    }
    return ;

  };

  // 获取当前问题
  const getCurrentQuestion = () => {
    const section = getCurrentSection();
    if (!section) return null;
    if (!section.children) return null;
    return section.children[currentQuestionIndex];
  };

  // 渲染问题内容
  const renderQuestion = (question:API.Questions) => {
    if (!question){
      return null;
    }
    switch (question.type) {
      case 'single_choice':
        return renderSingleChoiceQuestion(question);
      case 'multiple_choice':
        return renderMultipleChoiceQuestion(question);
      case 'text':
        return renderTextQuestion(question);
      case 'h2':
        return '';
      case 'page':
        return '';
      case 'number':
        return numberDigit(question);
      case 'date':
        return datePicker(question);
      case 'rate':
        return '';



      default:
        return null;
    }
  };
  // 时间题
  const numberDigit = (question:API.Questions) => {
    if (!question.options){
      return ;
    }
    // 从本地状态恢复答案
    const initialValue = answers[question.id]?.[`question${question.id}`];
    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请输入' }]}
        initialValue={initialValue}
      >
      <ProFormDigit
        name={`question${question.id}`}
        label={question.content}
        min={1}
        max={10}
        fieldProps={{ precision: 0 }}
        placeholder="请输入"
        addonAfter="元"
        style={{ width: 40 }}
      />
      </Form.Item>);
  }


  // 时间题
  const datePicker = (question:API.Questions) => {
    if (!question.options){
      return ;
    }
    // 从本地状态恢复答案
    const initialValue = answers[question.id]?.[`question${question.id}`];
    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请填写时间' }]}
        initialValue={initialValue}
      >
      <ProFormDatePicker
      name={`question${question.id}`}
      label={question.content}
      />
    </Form.Item>);
  }


  // 渲染单选题
  const renderSingleChoiceQuestion = (question:API.Questions) => {
   if (!question.options){
     return ;
   }
    // 从本地状态恢复答案
    const initialValue = answers[question.id]?.[`question${question.id}`];
    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请选择一个选项' }]}
        initialValue={initialValue}
      >
        <Radio.Group>
          {question.options.map(option => (
            <Radio key='content' value='content'>
              {option.content}
            </Radio>
          ))}
        </Radio.Group>
      </Form.Item>
    );
  };

  // 渲染多选题
  const renderMultipleChoiceQuestion = (question:API.Questions) => {

    if (!question.options){
      return ;
    }
    // 从本地状态恢复答案
    const initialValue = answers[question.id]?.[`question${question.id}`];
    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请选择至少一个选项' }]}
        initialValue={initialValue}
      >
        <Checkbox.Group>
          {question.options.map(option => (
            <Checkbox  key='content' value='content'>
              {option.content}
            </Checkbox>
          ))}
        </Checkbox.Group>
      </Form.Item>
    );
  };

  // 渲染文本题
  const renderTextQuestion = (question :API.Questions) => {
    if (!question.options){
      return ;
    }
    // 从本地状态恢复答案
    const initialValue = answers[question.id]?.[`question${question.id}`];
    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请输入您的意见' }]}
        initialValue={initialValue}
      >
        <Input.TextArea key='content' value='content' rows={4} placeholder="请详细描述您的需求和建议..." />
      </Form.Item>
    );
  };

  // 渲染导航按钮
  const renderNavigationButtons = () => {
    const section = getCurrentSection();
    if (!section) return null;
    const isFirstQuestion = currentQuestionIndex === 0;
    const isLastQuestion = currentQuestionIndex === section.children.length - 1;
    const isLastSection = currentSectionIndex === questions.length - 1;




    return (
      <div className="survey-buttons">

        <Button
          type="default"
          icon={<ArrowLeftOutlined />}
          disabled={isFirstQuestion}
          onClick={moveToPreviousQuestion}
          className="prev-button"
        >
          上一题
        </Button>
        <Button
          type="primary"
          onClick={() => form.submit()}
          loading={isSubmitting}
          className="next-button"
        >
          {isLastQuestion && isLastSection ? '提交问卷' : '下一步'}
        </Button>
      </div>
    );
  };

  // 渲染感谢页面
  const renderThankYou = () => {
    return (
      <div className="thank-you-container">
        <div className="thank-you-card">
          <QuestionCircleOutlined className="thank-you-icon" />
          <Title level={3}>感谢您参与调查！</Title>
          <Text>您的反馈对我们非常重要，我们将根据您的意见改进养老服务。</Text>
          <Button
            type="primary"
            onClick={() => navigate('/')}
            className="finish-button"
          >
            完成
          </Button>
        </div>
      </div>
    );
  };

  if (showThankYou) {
    return renderThankYou();
  }

const surveyContainer = () => {
  return (
    <div className="survey-container">
      <PageContainer
        loading={loading}
      >
        <ProCard
          className="survey-card"
          title={survey.title}
          bordered={false}
        >
      <ProCard
        className="survey-card"
         title={getCurrentSection()?getCurrentSection().content:''}
        extra={<Text type="secondary">第 {currentSectionIndex + 1}/{questions.length} 部分</Text>}
        bordered={false}
      >
        <ProForm form={form} onFinish={handleQuestionSubmit} layout="vertical">
          <div className="question-number">
            问题 {currentQuestionIndex + 1}/{getCurrentSection()?getCurrentSection().children.length:''}
          </div>

          {renderQuestion(getCurrentQuestion())}

          {renderNavigationButtons()}
        </ProForm>

      </ProCard>
      </ProCard>
      </PageContainer>
    </div>
  );
}

      return surveyContainer();


}
export default Survey;
