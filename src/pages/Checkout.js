import React, { useContext, useState } from 'react';
import { CartContext } from '../context/CartContext';
import { AuthContext } from '../context/AuthContext';
import { useNavigate } from 'react-router-dom';

const Checkout = () => {
    const { cart, total, clearCart } = useContext(CartContext);
    const { user } = useContext(AuthContext);
    const [form, setForm] = useState({ name: user ? user.username : '', email: user ? user.email : '', address: '' });
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    const onSubmit = async (e) => {
        e.preventDefault();
        if (cart.length === 0) { alert('Giỏ hàng rỗng'); return; }
        setLoading(true);
        try {
            const payload = {
                userId: user ? user.id || 1 : 0,
                date: new Date().toISOString(),
                products: cart.map(i => ({ productId: i.id, quantity: i.quantity })),
            };
            // Replace this with your new backend API call to create a cart
            const saved = { id: Date.now() }; // Simulate a saved cart with a temporary ID
            // For demo: also keep a local record
            const orders = JSON.parse(localStorage.getItem('orders') || '[]');
            orders.push({ id: saved.id || Date.now(), payload, total, customer: form });
            localStorage.setItem('orders', JSON.stringify(orders));
            clearCart();
            alert('Đặt hàng thành công!');
            navigate('/');
        } catch (err) {
            console.error(err);
            alert('Lỗi khi gửi đơn (demo)');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="container">
            <h2>Checkout</h2>
            <form onSubmit={onSubmit} className="checkout-form">
                <label>Họ tên
                    <input value={form.name} onChange={e => setForm({ ...form, name: e.target.value })} required />
                </label>
                <label>Email
                    <input type="email" value={form.email} onChange={e => setForm({ ...form, email: e.target.value })} required />
                </label>
                <label>Địa chỉ giao hàng
                    <textarea value={form.address} onChange={e => setForm({ ...form, address: e.target.value })} required />
                </label>
                <div className="checkout-summary">
                    <p>Tổng tiền: ${total.toFixed(2)}</p>
                    <button type="submit" disabled={loading}>{loading ? 'Đang xử lý...' : 'Xác nhận đặt hàng'}</button>
                </div>
            </form>
        </div>
    );
};

export default Checkout;
