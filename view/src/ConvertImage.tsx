import React, {RefObject} from "react";
import {Button, Select, Table, Tag} from "antd";
import {RightOutlined, UploadOutlined} from "@ant-design/icons";
import './ConvertImage.less';
import {HttpClient} from "./Http";
import qs from 'qs';

export class ConvertImage extends React.Component<any, any> {
    private readonly uploadAdd: RefObject<HTMLInputElement>

    constructor(props: any) {
        super(props);
        this.uploadAdd = React.createRef();
        this.state = {
            imgFile: '',
            columns: [
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
                                <Select.Option value="png">png</Select.Option>
                                <Select.Option value="jpg">jpg</Select.Option>
                                <Select.Option value="jpeg">jpeg</Select.Option>
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
            ],
            data: [
                {
                    key: '1',
                    name: '1.jpg',
                    status: '未上传',
                }
            ]
        }
    }

    render() {
        return <>
            <Button type="primary" danger icon={<UploadOutlined/>} size="large" className="choose-button"
                    onClick={() => this.uploadAdd.current?.click()}>
                选择文件
            </Button>
            <input type="file" ref={this.uploadAdd} style={{display: 'none'}} onChange={e => this.chooseFile(e)}/>

            <Table columns={this.state.columns} dataSource={this.state.data} className="table" pagination={false}/>

            <Button type="primary" icon={<RightOutlined/>} size="large" className="upload-button"
                    disabled={!this.state.imgFile}
                    onClick={() => this.start()}>
                开始转换
            </Button>
        </>;
    }

    private chooseFile(e: React.ChangeEvent<HTMLInputElement>) {
        this.setState({
            imgFile: e.target.files?.[0]
        });
    }

    private async start() {
        const res = await HttpClient.post("/upload", qs.stringify({
            'type': 'jpg',
            'file': this.state.imgFile
        }));
        console.log(res);
    }
}
