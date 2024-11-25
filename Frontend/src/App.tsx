import React, { useState, useEffect } from "react";
import axios from "axios";
import { Table, Button, Input, Form, Modal, message } from "antd";

const App = () => {
  type Item = {
    id: number;
    name: string;
  };

  const [items, setItems] =useState<Item[]>([]);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [editingItem, setEditingItem] =  useState<Item | null>(null);
  const [form] = Form.useForm();

  // API Base URL
  const API_URL = "http://localhost:3001/api/signup";
  const GET_URL = "http://localhost:3001/api/users";
  const UPDATE_URI = "http://localhost:3001/api/users";
  const DELETE_URI = "http://localhost:3001/api/users";

  // Fetch all items
  const fetchItems = async () => {
    try {
      const response = await axios.get(GET_URL);
      setItems(response.data);
    } catch (error) {
      console.error("Error fetching items:", error);
    }
  };

  // Add or Update Item
  const handleSubmit = async (values: {name: string , password: string , email :string, id: string}) => {
    try {
      if (editingItem) {
        // Update item
        await axios.put(`${UPDATE_URI}/${editingItem.id}`, values);
        message.success("Item updated successfully");
      } else {
        // Add new item
        await axios.post(API_URL, values);
        message.success("Item added successfully");
      }
      form.resetFields();
      setEditingItem(null);
      setIsModalVisible(false);
      fetchItems();
    } catch (error) {
      console.error("Error saving item:", error);
      message.error("Error saving item");
    }
  };

  // Delete Item
  const handleDelete = async (id: number) => {
    try {
      await axios.delete(`${DELETE_URI}/${id}`);
      message.success("Item deleted successfully");
      fetchItems();
    } catch (error) {
      console.error("Error deleting item:", error);
      message.error("Error deleting item");
    }
  };

  // Show Modal for Add/Edit
  const showModal = (item: Item |null= null) => {
    console.log(item)
    setEditingItem(item);
    setIsModalVisible(true);
    if (item) {
      form.setFieldsValue(item);
    } else {
      form.resetFields();
    }
  };

  // Close Modal
  const handleCancel = () => {
    setEditingItem(null);
    setIsModalVisible(false);
  };

  useEffect(() => {
    fetchItems();
  }, []);

  return (
    <div style={{ padding: 200 }}>
      <h1>Go CRUD App</h1>
      <Button type="primary" onClick={() => showModal()}>
        Add Item
      </Button>
      <Table
        dataSource={items}
        rowKey="id"
        columns={[
          {
            title: "ID",
            dataIndex: "id",
            key: "id",
          },
          {
            title: "Name",
            dataIndex: "name",
            key: "name",
          },
          {
            title: "Email",
            dataIndex: "email",
            key: "email",
          },
          {
            title: "Actions",
            render: (text, record: Item) => (
              <>
                <Button type="link" onClick={() => showModal(record)}>
                  Edit
                </Button>
                <Button type="link" danger onClick={() => handleDelete(record.id)}>
                  Delete
                </Button>
              </>
            ),
          },
        ]}
        style={{ marginTop: 20 }}
      />
      <Modal
        title={editingItem ? "Edit Item" : "Add Item"}
        visible={isModalVisible}
        onCancel={handleCancel}
        footer={null}
      >
        <Form form={form} onFinish={handleSubmit} layout="vertical">
        <Form.Item
            label="Name"
            name="name"
            rules={[{ required: true, message: "Please enter a name" }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Email"
            name="email"
            rules={[{ required: true, message: "Please enter a email" }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Password"
            name="password"
            rules={[{ required: true, message: "Please enter a password" }]}
          >
            <Input />
          </Form.Item>
          <Button type="primary" htmlType="submit">
            {editingItem ? "Update" : "Add"}
          </Button>
        </Form>
      </Modal>
    </div>
  );
};

export default App;
