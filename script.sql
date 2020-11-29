create table ipasgo_pedidos
(
    id                       serial not null
        constraint ipasgo_pedidos_pk
            primary key,
    ds_beneficiario          varchar(255),
    ds_cpf                   varchar(255),
    ds_guia_prestador        varchar(255),
    cd_autorizacao           integer,
    ds_senha_autorizacao     varchar(255),
    dt_autorizacao           varchar(255),
    ds_glosa                 varchar(255),
    cd_glosa                 varchar(255),
    dt_expiracao_autorizacao varchar(255),
    ds_tipo_guia             varchar(255),
    ds_crm                   varchar(255),
    ds_medico                varchar(255),
    created_at               timestamp default now(),
    updated_at               timestamp,
    nr_carteira              integer
);

alter table ipasgo_pedidos
    owner to dicomvix;

create table ipasgo_exames
(
    id                serial not null
        constraint ipasgo_exames_pk
            primary key,
    ds_guia_prestador varchar(255),
    cd_procedimento   integer,
    ds_procedimento   varchar(255),
    qtd_solicitada    integer,
    qtd_autorizada    integer,
    cd_status         varchar(255),
    ds_glosa          varchar(255),
    cd_glosa          varchar(255),
    created_at        timestamp default now(),
    updated_at        timestamp
);

alter table ipasgo_exames
    owner to dicomvix;

