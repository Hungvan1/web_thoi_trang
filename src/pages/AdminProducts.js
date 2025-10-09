import React, { useEffect, useState } from 'react';

const AdminProducts = () => {
    const [products, setProducts] = useState([]);
    const [editing, setEditing] = useState(null);
    const [form, setForm] = useState({ title: '', price: '', description: '', image: '' });

    useEffect(() => {
        // Replace with your new backend API call to fetch products
        setProducts([]); // Initialize with an empty array
    }, []);

    const startEdit = (p) => {
        setEditing(p.id);
        setForm({ title: p.title, price: p.price, description: p.description, image: p.image });
    };

    const saveEdit = () => {
        setProducts(prev => prev.map(p => p.id === editing ? { ...p, ...form } : p));
        setEditing(null);
    };

    const deleteProd = (id) => {
        if (window.confirm('Xóa sản phẩm này?')) {
            setProducts(prev => prev.filter(p => p.id !== id));
        }
    };

    const addNew = () => {
        const newProd = { ...form, id: Date.now(), price: Number(form.price) };
        setProducts(prev => [...prev, newProd]);
        setForm({ title: '', price: '', description: '', image: '' });
    };

    return (
        <div className="container">
            <h2>Quản lý sản phẩm (Demo CRUD)</h2>

            <div className="admin-form">
                <input placeholder="Tên sản phẩm" value={form.title} onChange={e => setForm({ ...form, title: e.target.value })} />
                <input placeholder="Giá" type="number" value={form.price} onChange={e => setForm({ ...form, price: e.target.value })} />
                <input placeholder="Ảnh URL" value={form.image} onChange={e => setForm({ ...form, image: e.target.value })} />
                <textarea placeholder="Mô tả" value={form.description} onChange={e => setForm({ ...form, description: e.target.value })} />
                {editing ? (
                    <button onClick={saveEdit}>Lưu chỉnh sửa</button>
                ) : (
                    <button onClick={addNew}>Thêm sản phẩm mới</button>
                )}
            </div>

            <table className="admin-table">
                <thead>
                    <tr><th>ID</th><th>Ảnh</th><th>Tên</th><th>Giá</th><th>Thao tác</th></tr>
                </thead>
                <tbody>
                    {products.map(p => (
                        <tr key={p.id}>
                            <td>{p.id}</td>
                            <td><img src={p.image} alt="" style={{ width: 50 }} /></td>
                            <td>{p.title}</td>
                            <td>${p.price}</td>
                            <td>
                                <button onClick={() => startEdit(p)}>Sửa</button>
                                <button onClick={() => deleteProd(p.id)}>Xóa</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default AdminProducts;
