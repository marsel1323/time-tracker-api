--
-- PostgreSQL database dump
--

-- Dumped from database version 12.6
-- Dumped by pg_dump version 12.6

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
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: time_tracker
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    name text NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.categories OWNER TO time_tracker;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: time_tracker
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO time_tracker;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: time_tracker
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: goal_statistics; Type: TABLE; Schema: public; Owner: time_tracker
--

CREATE TABLE public.goal_statistics (
    id bigint NOT NULL,
    milliseconds bigint NOT NULL,
    goal_id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.goal_statistics OWNER TO time_tracker;

--
-- Name: goal_statistics_id_seq; Type: SEQUENCE; Schema: public; Owner: time_tracker
--

CREATE SEQUENCE public.goal_statistics_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.goal_statistics_id_seq OWNER TO time_tracker;

--
-- Name: goal_statistics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: time_tracker
--

ALTER SEQUENCE public.goal_statistics_id_seq OWNED BY public.goal_statistics.id;


--
-- Name: goals; Type: TABLE; Schema: public; Owner: time_tracker
--

CREATE TABLE public.goals (
    id bigint NOT NULL,
    name text NOT NULL,
    "time" bigint DEFAULT 0 NOT NULL,
    category_id bigint NOT NULL,
    active boolean DEFAULT true NOT NULL,
    less_is_better boolean DEFAULT false NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.goals OWNER TO time_tracker;

--
-- Name: goals_id_seq; Type: SEQUENCE; Schema: public; Owner: time_tracker
--

CREATE SEQUENCE public.goals_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.goals_id_seq OWNER TO time_tracker;

--
-- Name: goals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: time_tracker
--

ALTER SEQUENCE public.goals_id_seq OWNED BY public.goals.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: marsel
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO marsel;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: time_tracker
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO time_tracker;

--
-- Name: stats; Type: TABLE; Schema: public; Owner: time_tracker
--

CREATE TABLE public.stats (
    id bigint NOT NULL,
    milliseconds bigint NOT NULL,
    task_id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.stats OWNER TO time_tracker;

--
-- Name: stats_id_seq; Type: SEQUENCE; Schema: public; Owner: time_tracker
--

CREATE SEQUENCE public.stats_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.stats_id_seq OWNER TO time_tracker;

--
-- Name: stats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: time_tracker
--

ALTER SEQUENCE public.stats_id_seq OWNED BY public.stats.id;


--
-- Name: task; Type: TABLE; Schema: public; Owner: time_tracker
--

CREATE TABLE public.task (
    id bigint NOT NULL,
    name text NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    done boolean DEFAULT false
);


ALTER TABLE public.task OWNER TO time_tracker;

--
-- Name: task_id_seq; Type: SEQUENCE; Schema: public; Owner: time_tracker
--

CREATE SEQUENCE public.task_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_id_seq OWNER TO time_tracker;

--
-- Name: task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: time_tracker
--

ALTER SEQUENCE public.task_id_seq OWNED BY public.task.id;


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: goal_statistics id; Type: DEFAULT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.goal_statistics ALTER COLUMN id SET DEFAULT nextval('public.goal_statistics_id_seq'::regclass);


--
-- Name: goals id; Type: DEFAULT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.goals ALTER COLUMN id SET DEFAULT nextval('public.goals_id_seq'::regclass);


--
-- Name: stats id; Type: DEFAULT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.stats ALTER COLUMN id SET DEFAULT nextval('public.stats_id_seq'::regclass);


--
-- Name: task id; Type: DEFAULT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.task ALTER COLUMN id SET DEFAULT nextval('public.task_id_seq'::regclass);


--
-- Name: categories categories_name_key; Type: CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_name_key UNIQUE (name);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: goal_statistics goal_statistics_pkey; Type: CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.goal_statistics
    ADD CONSTRAINT goal_statistics_pkey PRIMARY KEY (id);


--
-- Name: goals goals_name_key; Type: CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.goals
    ADD CONSTRAINT goals_name_key UNIQUE (name);


--
-- Name: goals goals_pkey; Type: CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.goals
    ADD CONSTRAINT goals_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: stats stats_pkey; Type: CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.stats
    ADD CONSTRAINT stats_pkey PRIMARY KEY (id);


--
-- Name: task task_pkey; Type: CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: marsel
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: goal_statistics goal_statistics_goal_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.goal_statistics
    ADD CONSTRAINT goal_statistics_goal_id_fkey FOREIGN KEY (goal_id) REFERENCES public.goals(id) ON DELETE CASCADE;


--
-- Name: goals goals_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.goals
    ADD CONSTRAINT goals_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE CASCADE;


--
-- Name: stats stats_task_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: time_tracker
--

ALTER TABLE ONLY public.stats
    ADD CONSTRAINT stats_task_id_fkey FOREIGN KEY (task_id) REFERENCES public.task(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

