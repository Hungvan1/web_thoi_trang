import React, { useState, useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import { useNavigate } from 'react-router-dom';

const Register = () => {
    const [form, setForm] = useState({ username: '', email: '', password: '' });
    const { loginLocal } = useContext(AuthContext);
    const navigate = useNavigate();

    const onSubmit = async (e) => {
        e.preventDefault();
        try {
            const payload = {
                email: form.email,
                username: form.username,
                password: form.password,
                name: { firstname: form.username, lastname: '' },
                address: { city: '', street: '', number: 0, zipcode: '' },
                phone: ''
            };
            const saved = { id: Date.now(), username: form.username, email: form.email };
            loginLocal({ id: saved.id, username: saved.username, email: saved.email });
            alert('Đăng ký thành công!');
            navigate('/');
        } catch (err) {
            console.error(err);
            alert('Lỗi khi đăng ký (demo)');
        }
    };

    return (
        <div className="container auth-page">
            <h2>Đăng ký</h2>
            <form onSubmit={onSubmit} className="auth-form">
                <label>Username<input required value={form.username} onChange={e => setForm({ ...form, username: e.target.value })} /></label>
                <label>Email<input type="email" required value={form.email} onChange={e => setForm({ ...form, email: e.target.value })} /></label>
                <label>Password<input type="password" required value={form.password} onChange={e => setForm({ ...form, password: e.target.value })} /></label>
                <button type="submit">Đăng ký</button>
            </form>
        </div>
    );
};

export default Register;
