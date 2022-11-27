INSERT INTO defi_data.e_job
(id, create_time, update_time, code, cron, handler, handler_param, name, notify_emails, remark, status, create_user_id, update_user_id)
VALUES(1, '2021-02-04 08:11:14', NULL, 'cmp_liq_job', '*/15 * * * * ?', 'io.alpha.admin.modules.compound.job.CompoundLiquidationJob', NULL, 'cmp_liq_job', NULL, 'Compound Healthy Data Fetch Job, every 15 seconds', 1, 1, NULL);
