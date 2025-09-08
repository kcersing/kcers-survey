import { ProTable } from '@ant-design/pro-components';
import { Card, Descriptions, Menu, Modal } from 'antd';
import React, { useState, useEffect } from 'react';
import { getSurvey, treeQuestion, getHeatmap, questionBasicData, getSurveyStatistics } from '@/services/survey';
import { useParams } from "react-router";
import { DemoCustomColor } from '@/pages/survey/statistics/components/custom-color';
import { DemoRose } from '@/pages/survey/statistics/components/donut-rose';
import { DemoPie } from '@/pages/survey/statistics/components/spider-label';
import { HeatMap } from "@/pages/survey/statistics/components/heatmap";
const columns = [
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
    const [survey, setSurvey] = useState({});
    const [surveyData, setSurveyData] = useState([]);
    const [questionBasic, setQuestionBasic] = useState([]);
    const [menuItems, setMenuItems] = useState([]);
    let params = useParams();
    const surveyId = parseInt(params.id);
    useEffect(() => {
        const getSurveyInfo = async () => {
            const surveyData = await getSurvey({ id: surveyId });
            setSurvey(surveyData.data);
        };
        getSurveyInfo();
        const getSurveyStatis = async () => {
            const getSurveyStatisticsData = await getSurveyStatistics({ id: surveyId });
            setSurveyData(getSurveyStatisticsData.data);
        };
        getSurveyStatis();
    }, []);
    useEffect(() => {
        const fetchMenuItems = async () => {
            try {
                const response = await treeQuestion({ surveyId: surveyId }); // 根据实际情况传入参数
                const items = convertToMenuItems(response.data);
                console.log(response.data);
                setMenuItems(items);
            }
            catch (error) {
                console.error('获取菜单数据失败:', error);
            }
        };
        fetchMenuItems();
    }, []);
    const convertToMenuItems = (data) => {
        return data.map(item => ({
            key: item.value,
            label: item.title,
            title: item.title,
            // 如果有子节点，递归转换
            children: item.children ? convertToMenuItems(item.children) : undefined,
        }));
    };
    const [openheatmap, setOpenheatmap] = useState(false);
    const [openheatmapdata, setOpenheatmapdata] = useState(false);
    const showLoading = () => {
        const getOpenheatmap = async () => {
            const openheatmapData = await getHeatmap({ id: surveyId });
            setOpenheatmapdata(openheatmapData.data);
        };
        getOpenheatmap();
        setOpenheatmap(true);
    };
    console.log(surveyData);
    return (<>
      <Card>
        <ProTable columns={columns} rowKey="key" pagination={{
            showSizeChanger: true,
        }} tableRender={(_, dom) => (<div style={{
                display: 'flex',
                width: '100%',
            }}>
              <Menu onSelect={(e) => setKey(e.key)} style={{ width: 256 }} defaultSelectedKeys={['1']} defaultOpenKeys={['sub1']} mode="inline" items={menuItems}/>
              <div style={{
                flex: 1,
            }}>
                {dom}
              </div>
            </div>)} tableExtraRender={(_, data) => (<Card>
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


              {questionBasic ? <DemoCustomColor data={questionBasic}/> : null}
              {questionBasic ? <DemoRose data={questionBasic}/> : null}
              {/*{questionBasic?  <Demobase data={questionBasic}/>:null}*/}
              {questionBasic ? <DemoPie data={questionBasic}/> : null}


            </Card>)} params={{
            key,
        }} request={async () => {
            console.log(key);
            const getQuestionBasicData = await questionBasicData({ id: parseInt(key) });
            setQuestionBasic(getQuestionBasicData.data.data);
            // const ans = await getQuestionAnswersList({ id: parseInt(key) });
            return {
                success: true,
                data: getQuestionBasicData.data.data,
            };
        }} dateFormatter="string"/>
      </Card>

      <Modal title={<p>地图</p>} centered width={{
            xs: '90%',
            sm: '80%',
            md: '70%',
            lg: '70%',
            xl: '70%',
            xxl: '80%',
        }} footer={<></>} open={openheatmap} onCancel={() => setOpenheatmap(false)}>
        <HeatMap data={openheatmapdata}/>
      </Modal>


    </>);
};
//# sourceMappingURL=index.jsx.map
