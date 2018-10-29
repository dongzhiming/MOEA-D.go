name_func = {'UF1';'UF2';'UF3';'UF4';'UF5';'UF6';'UF7';'UF8';'UF9';'UF10'; ...
    'LZ09_F1';'LZ09_F2';'LZ09_F3';'LZ09_F4';'LZ09_F5';'LZ09_F6';'LZ09_F7';'LZ09_F8';'LZ09_F9';
    'ZDT1';'ZDT2';'ZDT3';'ZDT4';'ZDT6'; ...
    'DTLZ1';'DTLZ2';'DTLZ3';'DTLZ4';'DTLZ5';'DTLZ6';'DTLZ7';};

% name_func = {'UF1';'UF2';'UF3';'UF4';'UF5';'UF6';'UF7';'UF8';'UF9';'UF10';};

numOfProblems = length(name_func);
numOfRun = 30;                              % running 次数

igd_va = zeros(numOfRun + 4,numOfProblems);

for p=1:numOfProblems
    problem = name_func{p};
    [nobj, dim, pfname, points] = get_info(problem);
    PF = load(pfname);
    for r=1:numOfRun
        fname   = sprintf('../result/FUN/%s_%d.FUN', problem, r-1);
        PF0     = load(fname);
        ind     = cec09filter(PF0', points);
        PF1     = PF0(ind, :);
        igd_va(r,p) = IGD(PF', PF1');
    end
end


%% 统计分析
igd_va(numOfRun+1,:) = mean(igd_va(1:numOfRun,:));
igd_va(numOfRun+2,:) = std(igd_va(1:numOfRun,:));
igd_va(numOfRun+3,:) = min(igd_va(1:numOfRun,:));
igd_va(numOfRun+4,:) = max(igd_va(1:numOfRun,:));

%% 输出保存
IGD = fopen('MOEADDE_IGD','w');
IGDStatistical = fopen('MOEADDE_STS','w');

for i=1:numOfProblems
    fprintf(IGD,'%20.8s',name_func{i});
end
fprintf(IGD,'\n');

for i=1:numOfRun
    for j=1:numOfProblems+1
        if j==1
            fprintf(IGD,'%6.4s',int2str(i));
        else
            fprintf(IGD,'%20.12d',igd_va(i, j-1));
        end
    end
    fprintf(IGD,'\n');
end

Statistical = ['mean';'std ';'min ';'max ';];
for i=1:4
    fprintf(IGDStatistical,'%20.8s',Statistical(i,:));
end
fprintf(IGDStatistical,'\n');

for i=1:numOfProblems
    for j=1:5
        if j==1
            fprintf(IGDStatistical,'%10.8s',name_func{i});
        else
            fprintf(IGDStatistical,'%20.12d',igd_va(numOfRun+j-1, i));
        end
    end
    fprintf(IGDStatistical,'\n');
end

fclose(IGD);
fclose(IGDStatistical);
