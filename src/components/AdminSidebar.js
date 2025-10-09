import React from 'react';
import { Link } from 'react-router-dom';

const AdminSidebar = () => (
    <aside className="admin-sidebar">
        <h3>Admin</h3>
        <ul>
            <li><Link to="/admin">Tổng quan</Link></li>
            <li><Link to="/admin?view=products">Sản phẩm</Link></li>
            <li><Link to="/admin?view=orders">Đơn hàng</Link></li>
        </ul>
    </aside>
);

export default AdminSidebar;
