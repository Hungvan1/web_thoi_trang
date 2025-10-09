import React, { useEffect, useState } from 'react';
import { Bar } from 'react-chartjs-2';
import AdminSidebar from '../components/AdminSidebar';
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js';
ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const AdminDashboard = () => {
    const [carts, setCarts] = useState([]);
    const [products, setProducts] = useState([]);
    const [users, setUsers] = useState([]);
    const [ordersLocal, setOrdersLocal] = useState([]);

    useEffect(() => {
        // You'll need to replace these API calls with your new backend API calls.
        // For now, setting empty arrays to avoid errors.
        setCarts([]);
        setProducts([]);
        setUsers([]);
        setOrdersLocal(JSON.parse(localStorage.getItem('orders') || '[]'));
    }, []);

    // Build sales per product from carts (fake data)
    const salesMap = {};
    carts.forEach(c => {
        (c.products || []).forEach(p => {
            salesMap[p.productId] = (salesMap[p.productId] || 0) + p.quantity;
        });
    });
    // also include local orders
    ordersLocal.forEach(o => {
        (o.payload.products || []).forEach(p => {
            salesMap[p.productId] = (salesMap[p.productId] || 0) + p.quantity;
        });
    });

    const labels = [];
    const values = [];
    // Map product ids to titles for top N
    Object.keys(salesMap).slice(0, 10).forEach(pid => {
        const prod = products.find(p => String(p.id) === String(pid));
        labels.push(prod ? prod.title.slice(0, 20) : `ID ${pid}`);
        values.push(salesMap[pid]);
    });

    const data = {
        labels,
        datasets: [{ label: 'Số lượng bán', data: values, backgroundColor: 'rgba(54,162,235,0.6)' }]
    };

    return (
        <div className="admin-page container">
            <AdminSidebar />
            <main className="admin-main">
                <h2>Dashboard Admin</h2>
                <div className="admin-cards">
                    <div className="card">
                        <h3>Tổng đơn </h3>
                        <p>{carts.length}</p>
                    </div>
                    <div className="card">
                        <h3>Đơn lưu </h3>
                        <p>{ordersLocal.length}</p>
                    </div>
                    <div className="card">
                        <h3>Tổng người dùng</h3>
                        <p>{users.length}</p>
                    </div>
                </div>

                <div style={{ maxWidth: 800 }}>
                    <Bar data={data} />
                </div>

                <section>
                    <h3>Danh sách đơn </h3>
                    <ul>
                        {carts.slice(0, 10).map(c => (
                            <li key={c.id}>Order #{c.id} - products: {(c.products || []).length} - date: {c.date}</li>
                        ))}
                    </ul>
                </section>
            </main>
        </div>
    );
};

export default AdminDashboard;
