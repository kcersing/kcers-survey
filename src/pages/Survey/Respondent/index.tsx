import React, { useState, useEffect, useRef } from 'react';
import {

  Form,
  Radio,
  Checkbox,
  Input,
  Button,
  Typography,

} from 'antd';
import { QuestionCircleOutlined } from '@ant-design/icons';
import { useNavigate } from 'umi';
import {
  ProCard,

} from '@ant-design/pro-components';



import './survey.css';

const { Title, Text } = Typography;

// 问卷数据 - 这里使用您提供的数据
const surveyData = [
  {
    "content": "养老观念与意愿",
    "type": "h2",
    "options": null,
    "required": 1,
    "sort": 115,
    "id": 115,
    "children": [
      {
        "content": "您更认可以下哪个养老观念？",
        "type": "single_choice",
        "options": [
          {
            "serial": 0,
            "content": "养老靠自己",
            "inputs": 0
          },
          {
            "serial": 0,
            "content": "养老要靠子女",
            "inputs": 0
          },
          {
            "serial": 2,
            "content": "养老靠社会福利（政府政策）",
            "inputs": 2
          },
          {
            "serial": 3,
            "content": "养老靠社区支持",
            "inputs": 3
          }
        ],
        "required": 1,
        "sort": 116,
        "id": 116,
        "children": null,
        "jumpRules": {
          "questionId": 0,
          "answer": "",
          "nextQuestionId": 0,
          "operators": "",
          "valueNumber": 0
        },
        "surveyId": 1,
        "parentId": 115,
        "serial": "28"
      },
      {
        "content": "您更倾向哪种养老方式？",
        "type": "single_choice",
        "options": [
          {
            "serial": 0,
            "content": "居家养老",
            "inputs": 0
          },
          {
            "serial": 1,
            "content": "社区服务",
            "inputs": 1
          },
          {
            "serial": 2,
            "content": "机构养老（养老院/护理院）",
            "inputs": 2
          },
          {
            "serial": 3,
            "content": "互助养老（邻里/同龄群体互助）",
            "inputs": 3
          },
          {
            "serial": 4,
            "content": "旅居养老（季节性异地养老）",
            "inputs": 4
          }
        ],
        "required": 1,
        "sort": 117,
        "id": 117,
        "children": null,
        "jumpRules": {
          "questionId": 0,
          "answer": "",
          "nextQuestionId": 0,
          "operators": "",
          "valueNumber": 0
        },
        "surveyId": 1,
        "parentId": 115,
        "serial": "29"
      },
      {
        "content": "您愿意参加互助养老吗？",
        "type": "single_choice",
        "options": [
          {
            "serial": 0,
            "content": "非常愿意",
            "inputs": 0
          },
          {
            "serial": 1,
            "content": "比较愿意",
            "inputs": 1
          },
          {
            "serial": 2,
            "content": "不确定",
            "inputs": 2
          },
          {
            "serial": 3,
            "content": "不太愿意",
            "inputs": 3
          },
          {
            "serial": 4,
            "content": "完全不愿意",
            "inputs": 4
          }
        ],
        "required": 1,
        "sort": 118,
        "id": 118,
        "children": null,
        "jumpRules": {
          "questionId": 0,
          "answer": "",
          "nextQuestionId": 0,
          "operators": "",
          "valueNumber": 0
        },
        "surveyId": 1,
        "parentId": 115,
        "serial": "30"
      },
      {
        "content": "您是否愿意入住养老机构？",
        "type": "single_choice",
        "options": [
          {
            "serial": 0,
            "content": "非常愿意",
            "inputs": 0
          },
          {
            "serial": 1,
            "content": "愿意",
            "inputs": 1
          },
          {
            "serial": 2,
            "content": "不确定",
            "inputs": 2
          },
          {
            "serial": 3,
            "content": "不太愿意",
            "inputs": 3
          },
          {
            "serial": 4,
            "content": "完全不愿意",
            "inputs": 4
          }
        ],
        "required": 1,
        "sort": 119,
        "id": 119,
        "children": null,
        "jumpRules": {
          "questionId": 119,
          "answer": "3,4",
          "nextQuestionId": 152,
          "operators": "包含",
          "valueNumber": 0
        },
        "surveyId": 1,
        "parentId": 115,
        "serial": "31"
      },
      {
        "content": "您希望机构能够提供以下哪些服务？（多选）",
        "type": "multiple_choice",
        "options": [
          {
            "serial": 0,
            "content": "日常生活照料（饮食/清洁/穿衣）",
            "inputs": 0
          },
          {
            "serial": 1,
            "content": "医疗护理（基础诊疗/康复护理）",
            "inputs": 1
          },
          {
            "serial": 2,
            "content": "康复保健（理疗/健身指导）",
            "inputs": 2
          },
          {
            "serial": 3,
            "content": "文化娱乐（活动/课程/社交）",
            "inputs": 3
          },
          {
            "serial": 4,
            "content": "心理咨询（情绪疏导/认知支持）",
            "inputs": 4
          },
          {
            "serial": 5,
            "content": "紧急救援（24小时监护/应急响应）",
            "inputs": 5
          }
        ],
        "required": 1,
        "sort": 121,
        "id": 121,
        "children": null,
        "jumpRules": {
          "questionId": 0,
          "answer": "",
          "nextQuestionId": 0,
          "operators": "",
          "valueNumber": 0
        },
        "surveyId": 1,
        "parentId": 115,
        "serial": "33"
      }
    ],
    "jumpRules": {
      "questionId": 0,
      "answer": "",
      "nextQuestionId": 0,
      "operators": "",
      "valueNumber": 0
    },
    "surveyId": 1,
    "parentId": 0,
    "serial": ""
  },
  {
    "content": "改进建议",
    "type": "h2",
    "options": null,
    "required": 1,
    "sort": 0,
    "id": 152,
    "children": [
      {
        "content": "您觉得当前农村养老服务在哪些方面还不能满足您的实际需要？",
        "type": "text",
        "options": null,
        "required": 1,
        "sort": 0,
        "id": 154,
        "children": null,
        "jumpRules": {
          "questionId": 0,
          "answer": "",
          "nextQuestionId": 0,
          "operators": "",
          "valueNumber": 0
        },
        "surveyId": 1,
        "parentId": 152,
        "serial": "45"
      }
    ],
    "jumpRules": {
      "questionId": 0,
      "answer": "",
      "nextQuestionId": 0,
      "operators": "",
      "valueNumber": 0
    },
    "surveyId": 1,
    "parentId": 0,
    "serial": ""
  }
];

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

  // 初始化ref数组
  useEffect(() => {
    sectionRefs.current = Array(surveyData.length).fill(null);
  }, [surveyData.length]);

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
    for (let s = 0; s < surveyData.length; s++) {
      const section = surveyData[s];
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
    const section = surveyData[currentSectionIndex];
    if (!section.children) return;

    const nextQuestionIndex = currentQuestionIndex + 1;
    if (nextQuestionIndex < section.children.length) {
      setCurrentQuestionIndex(nextQuestionIndex);
    } else {
      // 本部分已完成，移动到下一部分
      const nextSectionIndex = currentSectionIndex + 1;
      if (nextSectionIndex < surveyData.length) {
        setCurrentSectionIndex(nextSectionIndex);
        setCurrentQuestionIndex(0);
      } else {
        // 所有问题已完成
        handleSurveySubmit();
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
    return surveyData[currentSectionIndex];
  };

  // 获取当前问题
  const getCurrentQuestion = () => {
    const section = getCurrentSection();
    if (!section.children) return null;
    return section.children[currentQuestionIndex];
  };

  // 渲染问题内容
  const renderQuestion = (question) => {
    switch (question.type) {
      case 'single_choice':
        return renderSingleChoiceQuestion(question);
      case 'multiple_choice':
        return renderMultipleChoiceQuestion(question);
      case 'text':
        return renderTextQuestion(question);
      default:
        return null;
    }
  };

  // 渲染单选题
  const renderSingleChoiceQuestion = (question) => {
    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请选择一个选项' }]}
      >
        <Radio.Group>
          {question.options.map(option => (
            <Radio key={option.serial} value={option.serial}>
              {option.content}
            </Radio>
          ))}
        </Radio.Group>
      </Form.Item>
    );
  };

  // 渲染多选题
  const renderMultipleChoiceQuestion = (question) => {
    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请选择至少一个选项' }]}
      >
        <Checkbox.Group>
          {question.options.map(option => (
            <Checkbox key={option.serial} value={option.serial}>
              {option.content}
            </Checkbox>
          ))}
        </Checkbox.Group>
      </Form.Item>
    );
  };

  // 渲染文本题
  const renderTextQuestion = (question) => {
    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请输入您的意见' }]}
      >
        <Input.TextArea rows={4} placeholder="请详细描述您的需求和建议..." />
      </Form.Item>
    );
  };

  // 渲染导航按钮
  const renderNavigationButtons = () => {
    const section = getCurrentSection();
    const isLastQuestion = currentQuestionIndex === section.children.length - 1;
    const isLastSection = currentSectionIndex === surveyData.length - 1;

    return (
      <div className="survey-buttons">
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

  return (
    <div className="survey-container">
      <ProCard
        className="survey-card"
        title={getCurrentSection().content}
        extra={<Text type="secondary">第 {currentSectionIndex + 1}/{surveyData.length} 部分</Text>}
        bordered={false}
      >
        <Form form={form} onFinish={handleQuestionSubmit} layout="vertical">
          <div className="question-number">
            问题 {currentQuestionIndex + 1}/{getCurrentSection().children.length}
          </div>

          {renderQuestion(getCurrentQuestion())}

          {renderNavigationButtons()}
        </Form>
      </ProCard>
    </div>
  );
};

export default Survey;
