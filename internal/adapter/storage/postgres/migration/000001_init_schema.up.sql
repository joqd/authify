--
-- PostgreSQL database dump
--

-- Dumped from database version 15.12 (Debian 15.12-1.pgdg120+1)
-- Dumped by pg_dump version 15.12

-- Started on 2025-04-24 22:29:49 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3447 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 223 (class 1259 OID 16603)
-- Name: access_tokens; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.access_tokens (
    token character varying(255) NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.access_tokens OWNER TO admin;

--
-- TOC entry 224 (class 1259 OID 16609)
-- Name: api_keys; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.api_keys (
    key character varying(255) NOT NULL,
    secret character varying(255) NOT NULL,
    owner text NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.api_keys OWNER TO admin;

--
-- TOC entry 225 (class 1259 OID 16617)
-- Name: email_verification_tokens; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.email_verification_tokens (
    token character varying(255) NOT NULL,
    user_id bigint NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.email_verification_tokens OWNER TO admin;

--
-- TOC entry 222 (class 1259 OID 16588)
-- Name: group_permissions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.group_permissions (
    group_model_id bigint NOT NULL,
    permission_model_id bigint NOT NULL
);


ALTER TABLE public.group_permissions OWNER TO admin;

--
-- TOC entry 220 (class 1259 OID 16565)
-- Name: groups; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.groups (
    id bigint NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.groups OWNER TO admin;

--
-- TOC entry 219 (class 1259 OID 16564)
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.groups_id_seq OWNER TO admin;

--
-- TOC entry 3448 (class 0 OID 0)
-- Dependencies: 219
-- Name: groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


--
-- TOC entry 226 (class 1259 OID 16623)
-- Name: password_reset_tokens; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.password_reset_tokens (
    token character varying(255) NOT NULL,
    user_id bigint NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.password_reset_tokens OWNER TO admin;

--
-- TOC entry 217 (class 1259 OID 16537)
-- Name: permissions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.permissions (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    code_name character varying(255) NOT NULL,
    description text
);


ALTER TABLE public.permissions OWNER TO admin;

--
-- TOC entry 216 (class 1259 OID 16536)
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.permissions_id_seq OWNER TO admin;

--
-- TOC entry 3449 (class 0 OID 0)
-- Dependencies: 216
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- TOC entry 227 (class 1259 OID 16629)
-- Name: refresh_tokens; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.refresh_tokens (
    token character varying(255) NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.refresh_tokens OWNER TO admin;

--
-- TOC entry 229 (class 1259 OID 16636)
-- Name: user_activities; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.user_activities (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    event text,
    created_at timestamp with time zone
);


ALTER TABLE public.user_activities OWNER TO admin;

--
-- TOC entry 228 (class 1259 OID 16635)
-- Name: user_activities_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.user_activities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_activities_id_seq OWNER TO admin;

--
-- TOC entry 3450 (class 0 OID 0)
-- Dependencies: 228
-- Name: user_activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.user_activities_id_seq OWNED BY public.user_activities.id;


--
-- TOC entry 221 (class 1259 OID 16573)
-- Name: user_groups; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.user_groups (
    user_model_id bigint NOT NULL,
    group_model_id bigint NOT NULL
);


ALTER TABLE public.user_groups OWNER TO admin;

--
-- TOC entry 218 (class 1259 OID 16549)
-- Name: user_permissions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.user_permissions (
    user_model_id bigint NOT NULL,
    permission_model_id bigint NOT NULL
);


ALTER TABLE public.user_permissions OWNER TO admin;

--
-- TOC entry 215 (class 1259 OID 16520)
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    username character varying(255) NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password_hash text NOT NULL,
    is_staff boolean DEFAULT false,
    is_active boolean DEFAULT true,
    is_superuser boolean DEFAULT false,
    is_email_verified boolean DEFAULT false,
    last_login timestamp without time zone,
    date_joined timestamp with time zone
);


ALTER TABLE public.users OWNER TO admin;

--
-- TOC entry 214 (class 1259 OID 16519)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO admin;

--
-- TOC entry 3451 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3252 (class 2604 OID 16568)
-- Name: groups id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


--
-- TOC entry 3251 (class 2604 OID 16540)
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- TOC entry 3253 (class 2604 OID 16639)
-- Name: user_activities id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_activities ALTER COLUMN id SET DEFAULT nextval('public.user_activities_id_seq'::regclass);


--
-- TOC entry 3246 (class 2604 OID 16523)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3277 (class 2606 OID 16607)
-- Name: access_tokens access_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.access_tokens
    ADD CONSTRAINT access_tokens_pkey PRIMARY KEY (token);


--
-- TOC entry 3280 (class 2606 OID 16615)
-- Name: api_keys api_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.api_keys
    ADD CONSTRAINT api_keys_pkey PRIMARY KEY (key);


--
-- TOC entry 3283 (class 2606 OID 16621)
-- Name: email_verification_tokens email_verification_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.email_verification_tokens
    ADD CONSTRAINT email_verification_tokens_pkey PRIMARY KEY (token);


--
-- TOC entry 3275 (class 2606 OID 16592)
-- Name: group_permissions group_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.group_permissions
    ADD CONSTRAINT group_permissions_pkey PRIMARY KEY (group_model_id, permission_model_id);


--
-- TOC entry 3269 (class 2606 OID 16570)
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- TOC entry 3287 (class 2606 OID 16627)
-- Name: password_reset_tokens password_reset_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.password_reset_tokens
    ADD CONSTRAINT password_reset_tokens_pkey PRIMARY KEY (token);


--
-- TOC entry 3261 (class 2606 OID 16544)
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- TOC entry 3290 (class 2606 OID 16633)
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (token);


--
-- TOC entry 3271 (class 2606 OID 16572)
-- Name: groups uni_groups_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT uni_groups_name UNIQUE (name);


--
-- TOC entry 3263 (class 2606 OID 16548)
-- Name: permissions uni_permissions_code_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT uni_permissions_code_name UNIQUE (code_name);


--
-- TOC entry 3265 (class 2606 OID 16546)
-- Name: permissions uni_permissions_name; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT uni_permissions_name UNIQUE (name);


--
-- TOC entry 3255 (class 2606 OID 16535)
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- TOC entry 3257 (class 2606 OID 16533)
-- Name: users uni_users_username; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_username UNIQUE (username);


--
-- TOC entry 3293 (class 2606 OID 16643)
-- Name: user_activities user_activities_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_activities
    ADD CONSTRAINT user_activities_pkey PRIMARY KEY (id);


--
-- TOC entry 3273 (class 2606 OID 16577)
-- Name: user_groups user_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_groups
    ADD CONSTRAINT user_groups_pkey PRIMARY KEY (user_model_id, group_model_id);


--
-- TOC entry 3267 (class 2606 OID 16553)
-- Name: user_permissions user_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT user_permissions_pkey PRIMARY KEY (user_model_id, permission_model_id);


--
-- TOC entry 3259 (class 2606 OID 16531)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3278 (class 1259 OID 16608)
-- Name: idx_access_tokens_user_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_access_tokens_user_id ON public.access_tokens USING btree (user_id);


--
-- TOC entry 3281 (class 1259 OID 16616)
-- Name: idx_api_keys_owner; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_api_keys_owner ON public.api_keys USING btree (owner);


--
-- TOC entry 3284 (class 1259 OID 16622)
-- Name: idx_email_verification_tokens_user_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_email_verification_tokens_user_id ON public.email_verification_tokens USING btree (user_id);


--
-- TOC entry 3285 (class 1259 OID 16628)
-- Name: idx_password_reset_tokens_user_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_password_reset_tokens_user_id ON public.password_reset_tokens USING btree (user_id);


--
-- TOC entry 3288 (class 1259 OID 16634)
-- Name: idx_refresh_tokens_user_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_refresh_tokens_user_id ON public.refresh_tokens USING btree (user_id);


--
-- TOC entry 3291 (class 1259 OID 16644)
-- Name: idx_user_activities_user_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_user_activities_user_id ON public.user_activities USING btree (user_id);


--
-- TOC entry 3298 (class 2606 OID 16593)
-- Name: group_permissions fk_group_permissions_group_model; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.group_permissions
    ADD CONSTRAINT fk_group_permissions_group_model FOREIGN KEY (group_model_id) REFERENCES public.groups(id) ON UPDATE CASCADE;


--
-- TOC entry 3299 (class 2606 OID 16598)
-- Name: group_permissions fk_group_permissions_permission_model; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.group_permissions
    ADD CONSTRAINT fk_group_permissions_permission_model FOREIGN KEY (permission_model_id) REFERENCES public.permissions(id) ON UPDATE CASCADE;


--
-- TOC entry 3296 (class 2606 OID 16583)
-- Name: user_groups fk_user_groups_group_model; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_groups
    ADD CONSTRAINT fk_user_groups_group_model FOREIGN KEY (group_model_id) REFERENCES public.groups(id) ON UPDATE CASCADE;


--
-- TOC entry 3297 (class 2606 OID 16578)
-- Name: user_groups fk_user_groups_user_model; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_groups
    ADD CONSTRAINT fk_user_groups_user_model FOREIGN KEY (user_model_id) REFERENCES public.users(id) ON UPDATE CASCADE;


--
-- TOC entry 3294 (class 2606 OID 16559)
-- Name: user_permissions fk_user_permissions_permission_model; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT fk_user_permissions_permission_model FOREIGN KEY (permission_model_id) REFERENCES public.permissions(id) ON UPDATE CASCADE;


--
-- TOC entry 3295 (class 2606 OID 16554)
-- Name: user_permissions fk_user_permissions_user_model; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT fk_user_permissions_user_model FOREIGN KEY (user_model_id) REFERENCES public.users(id) ON UPDATE CASCADE;


-- Completed on 2025-04-24 22:29:49 UTC

--
-- PostgreSQL database dump complete
--

