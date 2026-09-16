
import type { ProColumns } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import { Card, Descriptions, Menu, Modal,Button } from 'antd';
import React, { useState, useEffect, useRef } from 'react';
import {getSurvey, treeQuestion,getQuestionAnswersList,getHeatmap,questionBasicData,getSurveyStatistics} from '@/services/ant-design-pro/survey';
import {useParams} from "react-router";
import { DemoCustomColor } from '@/pages/survey/statistics/components/custom-color';
import { DemoRose } from '@/pages/survey/statistics/components/donut-rose';
import { Demobase } from '@/pages/survey/statistics/components/space-layer';
import { DemoPie } from '@/pages/survey/statistics/components/spider-label';
import {HeatMap} from "@/pages/survey/statistics/components/heatmap";
import {ChinaHeatMap} from "@/pages/survey/statistics/components/china-heatmap";
import {L7HeatMap} from "@/pages/survey/statistics/components/l7-heatmap";

const columns: ProColumns[] = [


{
  title: '回答',
      dataIndex: 'type',

},
{
  title: '数量',
      dataIndex: 'value',


},

  // {
  //   title: '回答',
  //   dataIndex: 'answer',
  //   render: (_, record) =>(record.answer?record.answer.map(v=>( <>{v}  {record.answerText?("："+record.answerText):""};</> ) ):""),
  // },

  // 回答内容:{item.answer.map(v=>{ return (<a>{v}；</a>)} )}

  // {
  //   title: '创建时间',
  //   key: 'since2',
  //   dataIndex: 'createdAt',
  //   valueType: 'date',
  // },
];

export default () => {
  const [key, setKey] = useState('1');
  const [survey, setSurvey] = useState<API.Survey>({});

  const [surveyData, setSurveyData] = useState([]);
  const [questionBasic, setQuestionBasic] = useState([]);

  const [menuItems, setMenuItems] = useState<Menu.ItemType[]>([]);
  const questionTypeMap = useRef<Record<string, string>>({});


  let params = useParams();
  const surveyId =parseInt(params.id)
  useEffect(() => {
    const getSurveyInfo = async () => {
      const surveyData = await getSurvey({id: surveyId})
      setSurvey(surveyData.data);
    }
    getSurveyInfo()

    const getSurveyStatis = async () => {
      const getSurveyStatisticsData = await getSurveyStatistics({id: surveyId})
      setSurveyData(getSurveyStatisticsData.data);
    }
    getSurveyStatis()



  }, []);

  useEffect(() => {
    const fetchMenuItems = async () => {
      try {
        const response = await treeQuestion({ surveyId: surveyId }); // 根据实际情况传入参数
        const items = convertToMenuItems(response.data);
        setMenuItems(items);
      } catch (error) {
        console.error('获取菜单数据失败:', error);
      }
    };

    fetchMenuItems();
  }, []);

  const convertToMenuItems = (data: any[]): Menu.ItemType[] => {
    return data.map(item => {
      questionTypeMap.current[item.value] = item.type || '';
      return {
        key: item.value,
        label: item.title,
        title: item.title,
        children: item.children ? convertToMenuItems(item.children) : undefined,
      };
    });
  };


  const [openheatmap, setOpenheatmap] = useState<boolean>(false);
  const [openheatmapdata, setOpenheatmapdata] = useState<any[] | null>(null);
  const showLoading = () => {
    setOpenheatmap(true);
    getHeatmap({id: surveyId}).then((res) => {
      setOpenheatmapdata(res.data || []);
    });
  };

  const [openchina, setOpenchina] = useState<boolean>(false);
  const [openchinadata, setOpenchinadata] = useState<any[] | null>(null);
  const showChinaLoading = () => {
    setOpenchina(true);
    getHeatmap({id: surveyId}).then((res) => {
      setOpenchinadata(res.data || []);
    });
  };

  const [openl7, setOpenl7] = useState<boolean>(false);
  const [openl7data, setOpenl7data] = useState<any[] | null>(null);
  const showL7Loading = () => {
    setOpenl7(true);
    getHeatmap({id: surveyId}).then((res) => {
      setOpenl7data(res.data || []);
    });
  };

  return (
    <>
      <Card>
        <ProTable<API.ResponseAnswers>
          columns={columns}
          rowKey="key"
          pagination={{
            showSizeChanger: true,
          }}
          tableRender={(_, dom) => (
            <div
              style={{
                display: 'flex',
                width: '100%',
              }}
            >
              <Menu
                onSelect={(e) => setKey(e.key as string)}
                style={{ width: 256 }}
                defaultSelectedKeys={['1']}
                defaultOpenKeys={['sub1']}
                mode="inline"
                items={menuItems}
              />
              <div
                style={{
                  flex: 1,
                }}
              >
                {dom}
              </div>
            </div>
          )}


          tableExtraRender={(_, data) => (
            <Card>
              <h3> {survey.title ? survey.title : ''}</h3>
              <Descriptions size="small" column={3}>
                <Descriptions.Item label="已完成问卷数量">{surveyData.count ? surveyData.count : ''}</Descriptions.Item>
                <Descriptions.Item label="共完成题目">{surveyData.answersCount ? surveyData.answersCount : ''}</Descriptions.Item>
                <Descriptions.Item label="受访人">{surveyData.respondentCount ? surveyData.respondentCount : ''}</Descriptions.Item>
                <Descriptions.Item label="调研员">{surveyData.researcherCount ? surveyData.researcherCount : ''}</Descriptions.Item>
                <Descriptions.Item label="走访村庄">{surveyData.villageCount ? surveyData.villageCount : ''}</Descriptions.Item>
              </Descriptions>
              <a type="primary" onClick={showLoading}>
                地图
              </a>
              <a type="primary" onClick={showChinaLoading} style={{ marginLeft: 16 }}>
                热力地图
              </a>
              <a type="primary" onClick={showL7Loading} style={{ marginLeft: 16 }}>
                L7 热力图
              </a>


              {questionBasic? <DemoCustomColor data={questionBasic}/>:null}
              {questionBasic?  <DemoRose data={questionBasic}/>:null}
              {/*{questionBasic?  <Demobase data={questionBasic}/>:null}*/}
              {questionBasic?  <DemoPie data={questionBasic}/>:null}


            </Card>
          )}
          params={{
            key,
          }}
          request={async () => {
              const qt = questionTypeMap.current[key] || '';
              const getQuestionBasicData = await questionBasicData({id: parseInt(key), type: qt === 'address' || qt === 'address_input' ? 'area' : ''})

              setQuestionBasic(getQuestionBasicData.data.data);

            return {
              success: true,
              data: getQuestionBasicData.data.data,
            };
          }}
          dateFormatter="string"
        />
      </Card>

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

      <Modal
        title={<p>热力地图</p>}
        width="100vw"
        style={{ top: 0, maxWidth: '100vw', paddingBottom: 0 }}
        styles={{ body: { height: 'calc(100vh - 120px)', padding: 0 } }}
        footer={<></>}
        open={openchina}
        onCancel={() => setOpenchina(false)}
      >
        {openchinadata ? <ChinaHeatMap data={openchinadata} /> : null}
      </Modal>

      <Modal
        title={<p>L7 热力图</p>}
        width="100vw"
        style={{ top: 0, maxWidth: '100vw', paddingBottom: 0 }}
        styles={{ body: { height: 'calc(100vh - 120px)', padding: 0 } }}
        footer={<></>}
        open={openl7}
        onCancel={() => setOpenl7(false)}
      >
        {openl7data ? <L7HeatMap data={openl7data} /> : null}
      </Modal>


    </>
  );
};
