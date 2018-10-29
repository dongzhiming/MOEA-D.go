pf = load('result/FUN/UF1_2.FUN');
tpf = load('resources/pareto_fronts/UF1.pf');
pf = pf';
tpf = tpf';

plot(pf(1,:),pf(2,:),'ro')
hold on
plot(tpf(1,:),tpf(2,:),'b.')