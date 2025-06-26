import React, { useState, useEffect, useRef } from 'react';
import {
  Card,
  Form,
  Radio,
  Checkbox,
  Input,
  Button,
  Typography,
  Alert,
  Space,
  Rate,
  DatePicker,
  Upload,
  message,
  Spin,
  InputNumber,
} from 'antd';
import {
  QuestionCircleOutlined,
  ArrowLeftOutlined,
  ArrowRightOutlined,
  PlusOutlined,
  PictureOutlined,
  FileOutlined,
} from '@ant-design/icons';
import { useNavigate } from 'umi';
import ProCard from '@ant-design/pro-card';

import './survey.css';
import moment from 'moment';

const { Title, Text } = Typography;
const { TextArea } = Input;
const { MonthPicker, RangePicker, WeekPicker } = DatePicker;

// 问卷数据 - 包含所有类型的问题
const surveyData = [
  {
    "content": "养老服务调查问卷",
    "type": "h2",
    "options": null,
    "required": 1,
    "sort": 1,
    "id": 1,
    "children": [
      {
        "content": "基本信息",
        "type": "h3",
        "options": null,
        "required": 1,
        "sort": 2,
        "id": 2,
        "children": [
          {
            "content": "您的年龄：",
            "type": "number",
            "options": null,
            "required": 1,
            "sort": 3,
            "id": 3,
            "children": null,
            "jumpRules": null,
            "surveyId": 1,
            "parentId": 2,
            "serial": "1",
            "valueNumber": 0
          },
          {
            "content": "您的出生日期：",
            "type": "date",
            "options": null,
            "required": 1,
            "sort": 4,
            "id": 4,
            "children": null,
            "jumpRules": null,
            "surveyId": 1,
            "parentId": 2,
            "serial": "2",
            "valueNumber": 0
          },
          {
            "content": "您的性别：",
            "type": "single_choice",
            "options": [
              { "serial": 0, "content": "男", "inputs": 0 },
              { "serial": 1, "content": "女", "inputs": 1 },
              { "serial": 2, "content": "不愿透露", "inputs": 2 }
            ],
            "required": 1,
            "sort": 5,
            "id": 5,
            "children": null,
            "jumpRules": null,
            "surveyId": 1,
            "parentId": 2,
            "serial": "3",
            "valueNumber": 0
          }
        ],
        "jumpRules": null,
        "surveyId": 1,
        "parentId": 1,
        "serial": "",
        "valueNumber": 0
      },
      {
        "content": "养老观念与意愿",
        "type": "h3",
        "options": null,
        "required": 1,
        "sort": 6,
        "id": 6,
        "children": [
          {
            "content": "您是否愿意入住养老机构？",
            "type": "single_choice",
            "options": [
              { "serial": 0, "content": "完全不愿意", "inputs": 0 },
              { "serial": 1, "content": "比较不愿意", "inputs": 1 },
              { "serial": 2, "content": "一般", "inputs": 2 },
              { "serial": 3, "content": "比较愿意", "inputs": 3 },
              { "serial": 4, "content": "非常愿意", "inputs": 4 }
            ],
            "required": 1,
            "sort": 7,
            "id": 7,
            "children": [
              {
                "content": "您希望选择哪种类型的养老机构？",
                "type": "single_choice",
                "options": [
                  { "serial": 0, "content": "公立养老院", "inputs": 0 },
                  { "serial": 1, "content": "私立高端养老院", "inputs": 1 },
                  { "serial": 2, "content": "医养结合型机构", "inputs": 2 }
                ],
                "required": false,
                "sort": 11,
                "id": 11,
                "children": null,
                "jumpRules": null,
                "surveyId": 1,
                "parentId": 7,
                "serial": "7",
                "valueNumber": 0
              }
            ],
            "jumpRules": [
              { "answer": "一般", "nextQuestionId": 9, "operators": "equals" },
              { "answer": "比较愿意", "nextQuestionId": 9, "operators": "equals" },
              { "answer": "非常愿意", "nextQuestionId": 10, "operators": "equals" }
            ],
            "surveyId": 1,
            "parentId": 6,
            "serial": "4",
            "valueNumber": 0
          },
          {
            "content": "您不愿意入住养老机构的原因是什么？",
            "type": "multiple_choice",
            "options": [
              { "serial": 0, "content": "担心别人议论", "inputs": 0 },
              { "serial": 1, "content": "费用高", "inputs": 1 },
              { "serial": 2, "content": "离家远", "inputs": 2 },
              { "serial": 3, "content": "服务差", "inputs": 3 },
              { "serial": 4, "content": "没有私密性", "inputs": 4 },
              { "serial": 5, "content": "害怕难适应", "inputs": 5 },
              { "serial": 6, "content": "其他（请说明）", "inputs": 6 }
            ],
            "required": 1,
            "sort": 8,
            "id": 8,
            "children": [
              {
                "content": "如果费用是主要原因，您认为合理的收费标准是多少？",
                "type": "number",
                "options": null,
                "required": false,
                "sort": 12,
                "id": 12,
                "children": null,
                "jumpRules": null,
                "surveyId": 1,
                "parentId": 8,
                "serial": "8",
                "valueNumber": 0
              }
            ],
            "jumpRules": null,
            "surveyId": 1,
            "parentId": 6,
            "serial": "5",
            "valueNumber": 0
          },
          {
            "content": "您希望机构能够提供以下哪些服务？",
            "type": "multiple_choice",
            "options": [
              { "serial": 0, "content": "日常生活照料", "inputs": 0 },
              { "serial": 1, "content": "医疗护理", "inputs": 1 },
              { "serial": 2, "content": "康复训练", "inputs": 2 },
              { "serial": 3, "content": "心理慰藉", "inputs": 3 },
              { "serial": 4, "content": "临终关怀", "inputs": 4 },
              { "serial": 5, "content": "其他（请说明）", "inputs": 5 }
            ],
            "required": 1,
            "sort": 9,
            "id": 9,
            "children": [
              {
                "content": "如果选择了医疗护理，您希望多久进行一次健康检查？",
                "type": "single_choice",
                "options": [
                  { "serial": 0, "content": "每天", "inputs": 0 },
                  { "serial": 1, "content": "每周", "inputs": 1 },
                  { "serial": 2, "content": "每月", "inputs": 2 },
                  { "serial": 3, "content": "每季度", "inputs": 3 }
                ],
                "required": false,
                "sort": 13,
                "id": 13,
                "children": null,
                "jumpRules": null,
                "surveyId": 1,
                "parentId": 9,
                "serial": "9",
                "valueNumber": 0
              }
            ],
            "jumpRules": null,
            "surveyId": 1,
            "parentId": 6,
            "serial": "6",
            "valueNumber": 0
          },
          {
            "content": "您每月最多能够承担的养老机构费用：（元）",
            "type": "number",
            "options": null,
            "required": 1,
            "sort": 10,
            "id": 10,
            "children": [
              {
                "content": "如果您的预算较高，是否考虑过高端养老社区？",
                "type": "single_choice",
                "options": [
                  { "serial": 0, "content": "是", "inputs": 0 },
                  { "serial": 1, "content": "否", "inputs": 1 }
                ],
                "required": false,
                "sort": 14,
                "id": 14,
                "children": null,
                "jumpRules": null,
                "surveyId": 1,
                "parentId": 10,
                "serial": "10",
                "valueNumber": 0
              }
            ],
            "jumpRules": null,
            "surveyId": 1,
            "parentId": 6,
            "serial": "7",
            "valueNumber": 0
          },
          {
            "content": "您对当前养老服务的满意度：",
            "type": "rate",
            "options": null,
            "required": 1,
            "sort": 15,
            "id": 15,
            "children": null,
            "jumpRules": null,
            "surveyId": 1,
            "parentId": 6,
            "serial": "11",
            "valueNumber": 0
          }
        ],
        "jumpRules": null,
        "surveyId": 1,
        "parentId": 1,
        "serial": "",
        "valueNumber": 0
      },
      {
        "content": "改进建议",
        "type": "h3",
        "options": null,
        "required": 1,
        "sort": 16,
        "id": 16,
        "children": [
          {
            "content": "您觉得当前农村养老服务在哪些方面还不能满足您的实际需要？",
            "type": "text",
            "options": null,
            "required": 1,
            "sort": 17,
            "id": 17,
            "children": [
              {
                "content": "您认为最需要优先改进的是？",
                "type": "single_choice",
                "options": [
                  { "serial": 0, "content": "基础设施建设", "inputs": 0 },
                  { "serial": 1, "content": "医疗资源配置", "inputs": 1 },
                  { "serial": 2, "content": "服务人员培训", "inputs": 2 },
                  { "serial": 3, "content": "政策宣传落实", "inputs": 3 }
                ],
                "required": false,
                "sort": 18,
                "id": 18,
                "children": null,
                "jumpRules": null,
                "surveyId": 1,
                "parentId": 17,
                "serial": "12",
                "valueNumber": 0
              },
              {
                "content": "请上传相关图片或文件（可选）",
                "type": "uploadImage",
                "options": null,
                "required": false,
                "sort": 19,
                "id": 19,
                "children": null,
                "jumpRules": null,
                "surveyId": 1,
                "parentId": 17,
                "serial": "13",
                "valueNumber": 0
              },
              {
                "content": "请上传相关文件（可选）",
                "type": "uploadFile",
                "options": null,
                "required": false,
                "sort": 20,
                "id": 20,
                "children": null,
                "jumpRules": null,
                "surveyId": 1,
                "parentId": 17,
                "serial": "14",
                "valueNumber": 0
              }
            ],
            "jumpRules": null,
            "surveyId": 1,
            "parentId": 16,
            "serial": "12",
            "valueNumber": 0
          }
        ],
        "jumpRules": null,
        "surveyId": 1,
        "parentId": 1,
        "serial": "",
        "valueNumber": 0
      }
    ],
    "jumpRules": null,
    "surveyId": 1,
    "parentId": 0,
    "serial": "",
    "valueNumber": 0
  }
];

// 问卷调查组件
const Survey = () => {
  const [form] = Form.useForm();
  const [currentPath, setCurrentPath] = useState([0, 0, 0]); // 初始路径指向第一个问题
  const [answers, setAnswers] = useState({}); // 保存所有答案
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [showThankYou, setShowThankYou] = useState(false);
  const [uploading, setUploading] = useState(false);
  const navigate = useNavigate();
  const firstRender = useRef(true); // 用于标记首次渲染

  // 获取当前问题
  const getCurrentQuestion = (data, path) => {
    let current = data;
    for (const index of path) {
      if (current && current.children && Array.isArray(current.children) && current.children[index]) {
        current = current.children[index];
      } else {
        return null;
      }
    }
    return current;
  };

  // 查找问题ID对应的路径
  const findPathById = (data, id, path = []) => {
    if (!data) return null;
    if (data.id === id) return path;

    if (data.children && Array.isArray(data.children)) {
      for (let i = 0; i < data.children.length; i++) {
        const result = findPathById(data.children[i], id, [...path, i]);
        if (result) return result;
      }
    }
    return null;
  };

  // 处理多规则跳题
  const handleJumpRules = (question, values) => {
    if (!question.jumpRules || !Array.isArray(question.jumpRules) || question.jumpRules.length === 0) return null;

    const selectedAnswer = values[`question${question.id}`];

    for (const rule of question.jumpRules) {
      // 根据操作符判断是否满足条件
      if (rule.operators === 'equals' && selectedAnswer + '' === rule.answer) {
        return findPathById(surveyData, rule.nextQuestionId);
      }
    }
    return null;
  };

  // 移动到下一个问题
  const moveToNextQuestion = () => {
    const currentQuestion = getCurrentQuestion(surveyData, currentPath);
    if (!currentQuestion) {
      handleSurveySubmit();
      return;
    }

    if (currentQuestion.children && Array.isArray(currentQuestion.children) && currentQuestion.children.length > 0) {
      // 有子问题，移动到第一个子问题
      setCurrentPath([...currentPath, 0]);
      return;
    }

    // 没有子问题，尝试向上一层级移动
    if (currentPath.length === 0) {
      // 已经是最后一个问题，提交问卷
      handleSurveySubmit();
      return;
    }

    // 复制当前路径并移除最后一个索引
    let newPath = [...currentPath];
    newPath.pop();

    // 找到下一个兄弟节点
    let nextSibling = findNextSibling(surveyData, newPath);
    if (nextSibling) {
      setCurrentPath(nextSibling);
      return;
    }

    // 如果没有下一个兄弟节点，继续向上查找
    while (newPath.length > 0) {
      nextSibling = findNextSibling(surveyData, newPath);
      if (nextSibling) {
        setCurrentPath(nextSibling);
        return;
      }
      newPath.pop();
    }

    // 所有问题已完成
    handleSurveySubmit();
  };

  // 查找下一个兄弟节点路径
  const findNextSibling = (data, path) => {
    let current = data;
    const pathCopy = [...path];

    // 找到当前节点
    for (let i = 0; i < pathCopy.length; i++) {
      if (!current.children || !Array.isArray(current.children) || !current.children[pathCopy[i]]) {
        return null;
      }
      current = current.children[pathCopy[i]];
    }

    // 获取当前层级的下一个索引
    const lastIndex = pathCopy.pop();
    if (current.children && Array.isArray(current.children) && current.children.length > lastIndex + 1) {
      return [...pathCopy, lastIndex + 1];
    }

    return null;
  };

  // 移动到上一个问题
  const moveToPreviousQuestion = () => {
    if (currentPath.length === 0) return; // 已经是第一个问题

    const newPath = [...currentPath];

    // 尝试查找当前节点的最后一个子节点
    let current = getCurrentQuestion(surveyData, newPath);
    if (current && current.children && Array.isArray(current.children) && current.children.length > 0) {
      // 找到最后一个子节点
      let childPath = [...newPath];
      while (current && current.children && Array.isArray(current.children) && current.children.length > 0) {
        childPath.push(current.children.length - 1);
        current = current.children[current.children.length - 1];
      }
      setCurrentPath(childPath);
      return;
    }

    // 没有子节点，查找上一个兄弟节点
    if (newPath.length > 0) {
      const parentPath = [...newPath];
      const lastIndex = parentPath.pop();

      if (lastIndex > 0) {
        // 有上一个兄弟节点
        parentPath.push(lastIndex - 1);
        setCurrentPath(parentPath);
        return;
      }
    }

    // 没有上一个兄弟节点，移动到父节点
    newPath.pop();
    if (newPath.length > 0) {
      setCurrentPath(newPath);
    }
  };

  // 处理问题提交
  const handleQuestionSubmit = (values) => {
    const currentQuestion = getCurrentQuestion(surveyData, currentPath);
    if (currentQuestion) {
      setAnswers(prev => ({
        ...prev,
        [currentQuestion.id]: values
      }));

      // 处理跳题规则
      const jumpPath = handleJumpRules(currentQuestion, values);
      if (jumpPath) {
        setCurrentPath(jumpPath);
        return;
      }
    }

    moveToNextQuestion();
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

  // 处理文件上传
  const handleUpload = (type) => {
    setUploading(true);
    // 模拟上传
    setTimeout(() => {
      setUploading(false);
      message.success('上传成功');
    }, 1500);
  };

  // 递归渲染问题
  const renderQuestion = (data, depth) => {
    if (data.type === 'h2' || data.type === 'h3') {
      return renderHeading(data, depth);
    }

    switch (data.type) {
      case 'single_choice':
        return renderSingleChoiceQuestion(data, depth);
      case 'multiple_choice':
        return renderMultipleChoiceQuestion(data, depth);
      case 'text':
        return renderTextQuestion(data, depth);
      case 'number':
        return renderNumberQuestion(data, depth);
      case 'date':
        return renderDateQuestion(data, depth);
      case 'rate':
        return renderRateQuestion(data, depth);
      case 'uploadImage':
        return renderUploadImageQuestion(data, depth);
      case 'uploadFile':
        return renderUploadFileQuestion(data, depth);
      default:
        return null;
    }
  };

  // 渲染标题
  const renderHeading = (data, depth) => {
    const headingComponent = {
      'h2': Typography.H2,
      'h3': Typography.H3
    }[data.type];

    return (
      <div className={`heading heading-depth-${depth}`}>
        <headingComponent>{data.content}</headingComponent>
      </div>
    );
  };

  // 渲染单选题
  const renderSingleChoiceQuestion = (question, depth) => {
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
            <Radio key={option.serial} value={option.serial}>
              {option.content}
            </Radio>
          ))}
        </Radio.Group>
      </Form.Item>
    );
  };

  // 渲染多选题
  const renderMultipleChoiceQuestion = (question, depth) => {
    const initialValue = answers[question.id]?.[`question${question.id}`] || [];

    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请选择至少一个选项' }]}
        initialValue={initialValue}
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
  const renderTextQuestion = (question, depth) => {
    const initialValue = answers[question.id]?.[`question${question.id}`];

    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请输入您的意见' }]}
        initialValue={initialValue}
      >
        <TextArea rows={4} placeholder="请详细描述您的需求和建议..." />
      </Form.Item>
    );
  };

  // 渲染数字题
  const renderNumberQuestion = (question, depth) => {
    const initialValue = answers[question.id]?.[`question${question.id}`];

    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请输入数字' }]}
        initialValue={initialValue}
      >
        <InputNumber placeholder="请输入数字" />
      </Form.Item>
    );
  };

  // 渲染日期题
  const renderDateQuestion = (question, depth) => {
    const initialValue = answers[question.id]?.[`question${question.id}`]
      ? moment(answers[question.id][`question${question.id}`])
      : null;

    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请选择日期' }]}
        initialValue={initialValue}
      >
        <DatePicker placeholder="请选择日期" />
      </Form.Item>
    );
  };

  // 渲染评分题
  const renderRateQuestion = (question, depth) => {
    const initialValue = answers[question.id]?.[`question${question.id}`];

    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请进行评分' }]}
        initialValue={initialValue}
      >
        <Rate tooltips={['非常不满意', '不满意', '一般', '满意', '非常满意']} />
      </Form.Item>
    );
  };

  // 渲染上传图片题
  const renderUploadImageQuestion = (question, depth) => {
    const initialValue = answers[question.id]?.[`question${question.id}`];

    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请上传图片' }]}
        initialValue={initialValue}
      >
        <Upload
          name="file"
          listType="picture"
          className="upload-image"
          onStart={() => handleUpload('image')}
          onProgress={() => {}}
          onSuccess={() => {}}
          onError={() => {
            message.error('上传失败');
            setUploading(false);
          }}
          disabled={uploading}
        >
          {!uploading ? <Button icon={<PlusOutlined />}>上传图片</Button> : (
            <Spin size="small" tip="上传中..."/>
          )}
        </Upload>
      </Form.Item>
    );
  };

  // 渲染上传文件题
  const renderUploadFileQuestion = (question, depth) => {
    const initialValue = answers[question.id]?.[`question${question.id}`];

    return (
      <Form.Item
        name={`question${question.id}`}
        label={question.content}
        rules={[{ required: question.required, message: '请上传文件' }]}
        initialValue={initialValue}
      >
        <Upload
          name="file"
          listType="file"
          className="upload-file"
          onStart={() => handleUpload('file')}
          onProgress={() => {}}
          onSuccess={() => {}}
          onError={() => {
            message.error('上传失败');
            setUploading(false);
          }}
          disabled={uploading}
        >
          {!uploading ? <Button icon={<PlusOutlined />}>上传文件</Button> : (
            <Spin size="small" tip="上传中..."/>
          )}
        </Upload>
      </Form.Item>
    );
  };

  // 渲染导航按钮
  const renderNavigationButtons = () => {
    const currentQuestion = getCurrentQuestion(surveyData, currentPath);
    const isFirstQuestion = currentPath.length === 0 || (currentPath.length === 3 && currentPath[2] === 0);
    const hasChildren = currentQuestion && currentQuestion.children && Array.isArray(currentQuestion.children) && currentQuestion.children.length > 0;

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
          loading={isSubmitting || uploading}
          className="next-button"
        >
          {hasChildren || currentPath.length < 3 ? '下一步' : '提交问卷'}
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

  // 格式化进度显示
  const formatProgress = () => {
    let totalQuestions = 0;
    let answeredQuestions = 0;

    const countQuestions = (data) => {
      if (!data) return;

      if (data.type !== 'h2' && data.type !== 'h3') {
        totalQuestions++;
        if (answers[data.id] !== undefined) {
          answeredQuestions++;
        }
      }

      if (data.children && Array.isArray(data.children)) {
        data.children.forEach(child => countQuestions(child));
      }
    };

    countQuestions(surveyData[0]);

    return `${answeredQuestions}/${totalQuestions}`;
  };

  // 防止首次渲染时的无限循环
  useEffect(() => {
    firstRender.current = false;
  }, []);

  if (showThankYou) {
    return renderThankYou();
  }

  // 首次渲染时不执行路径重置，避免无限循环
  const currentQuestion = firstRender.current
    ? getCurrentQuestion(surveyData, currentPath)
    : getCurrentQuestion(surveyData, currentPath);

  if (!currentQuestion && !firstRender.current) {
    // 如果找不到当前问题，重置到第一个问题
    setCurrentPath([0, 0, 0]);
    return (
      <div className="loading-message">
        <Spin tip="加载问卷中..." />
      </div>
    );
  }

  return (
    <div className="survey-container">
      <ProCard
        className="survey-card"
        title={currentQuestion?.content || '加载中...'}
        extra={firstRender.current ? null : <Text type="secondary">答题进度: {formatProgress()}</Text>}
        bordered={false}
      >
        <Form form={form} onFinish={handleQuestionSubmit} layout="vertical">
          {firstRender.current ? null : renderQuestion(currentQuestion, currentPath.length)}

          {firstRender.current ? null : (
            currentQuestion && currentQuestion.type !== 'h2' && currentQuestion.type !== 'h3' && renderNavigationButtons()
          )}

          {/* 递归渲染子问题 */}
          {firstRender.current ? null : (
            currentQuestion && currentQuestion.children && Array.isArray(currentQuestion.children) && currentQuestion.children.length > 0 && (
              <div className="children-questions">
                {currentQuestion.children.map((child, index) => (
                  <div key={index} style={{ paddingLeft: `${currentPath.length * 20}px` }}>
                    {renderQuestion(child, currentPath.length + 1)}
                  </div>
                ))}
              </div>
            )
          )}
        </Form>
      </ProCard>
    </div>
  );
};

export default Survey;
