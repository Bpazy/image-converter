import React, {FC} from 'react';
import './App.less';
import {Button, Layout, Select, Table, Tag} from 'antd';
import {UploadOutlined, RightOutlined} from '@ant-design/icons';
import {Content, Header} from "antd/es/layout/layout";

const {Option} = Select;

const columns = [
    {
        title: '文件名',
        dataIndex: 'name',
        key: 'name',
        render(text: string) {
            return <a href="/#">{text}</a>;
        },
    },
    {
        title: '目标格式',
        dataIndex: 'type',
        key: 'type',
        render() {
            return <>
                <Select defaultValue="png" style={{width: 120}}>
                    <Option value="png">png</Option>
                    <Option value="jpg">jpg</Option>
                    <Option value="jpeg">jpeg</Option>
                </Select>
            </>
        },
    },
    {
        title: '状态',
        key: 'status',
        dataIndex: 'status',
        render(status: string) {
            let color = status.length > 5 ? 'geekblue' : 'green';
            if (status === 'loser') {
                color = 'volcano';
            }
            return (
                <Tag color={color} key={status}>
                    {status.toUpperCase()}
                </Tag>
            );
        },
    }
];

const data = [
    {
        key: '1',
        name: '1.jpg',
        status: '未上传',
    }
];

const App: FC = () => (
    <Layout>
        <Header>
            <div className="logo"/>
            <h1 className="page-title">在线图像文件转换器</h1>
        </Header>
        <Content className="content">
            <Button type="primary" danger icon={<UploadOutlined />} size="large" className="choose-button">
                选择文件
            </Button>
            <input type="file" id="upload-add" style={{display: 'none'}}/>

            <Table columns={columns} dataSource={data} className="table" pagination={false}/>

            <Button type="primary" icon={<RightOutlined />} size="large" className="upload-button">
                开始转换
            </Button>
        </Content>
    </Layout>
);

export default App;
