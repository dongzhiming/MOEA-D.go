function [nobj, dim, pfname, points] = get_info(problem)

%%
    if strcmp(problem, 'ZDT1') || strcmp(problem, 'ZDT2') || strcmp(problem, 'ZDT3') || strcmp(problem, 'ZDT4') || strcmp(problem, 'ZDT6')
        nobj = 2; dim = 10;
    elseif strcmp(problem, 'DTLZ1')
        nobj = 3; dim = 7;
    elseif strcmp(problem, 'DTLZ2') || strcmp(problem, 'DTLZ3') || strcmp(problem, 'DTLZ4') || strcmp(problem, 'DTLZ5') || strcmp(problem, 'DTLZ6')
        nobj = 3; dim = 12;
    elseif strcmp(problem, 'DTLZ7')
        nobj = 3; dim = 22;
    elseif strcmp(problem, 'UF1') || strcmp(problem, 'UF2') || strcmp(problem, 'UF3') || strcmp(problem, 'UF4') || strcmp(problem, 'UF5') || strcmp(problem, 'UF6') || strcmp(problem, 'UF7')
        nobj = 2; dim = 30;
    elseif strcmp(problem, 'UF8') || strcmp(problem, 'UF9') || strcmp(problem, 'UF10')
        nobj = 3; dim = 30;
    elseif strcmp(problem, 'LZ09_F1') || strcmp(problem, 'LZ09_F2') || strcmp(problem, 'LZ09_F3') || strcmp(problem, 'LZ09_F4') || strcmp(problem, 'LZ09_F5')
        nobj = 2; dim = 30;
    elseif strcmp(problem, 'LZ09_F6')
        nobj = 3; dim = 10;
    elseif strcmp(problem, 'LZ09_F7') || strcmp(problem, 'LZ09_F8')
        nobj = 2; dim = 10;
    elseif strcmp(problem, 'LZ09_F9')
        nobj = 2; dim = 30;
    else
        error('Undefined test problem name');
    end
    
    %%
    if (strcmp(problem,'DTLZ1')==1 || strcmp(problem,'DTLZ2')==1 || strcmp(problem,'DTLZ3')==1 || strcmp(problem,'DTLZ4')==1 || ...
        strcmp(problem,'DTLZ5')==1 || strcmp(problem,'DTLZ6')==1 || strcmp(problem,'DTLZ7')==1) && nobj == 3

        pfname = sprintf('../resources/pareto_fronts/%s.3D.pf', problem);
    elseif (strcmp(problem,'DTLZ1')==1 || strcmp(problem,'DTLZ2')==1 || strcmp(problem,'DTLZ3')==1 || strcmp(problem,'DTLZ4')==1 || ...
        strcmp(problem,'DTLZ5')==1 || strcmp(problem,'DTLZ6')==1 || strcmp(problem,'DTLZ7')==1) && nobj == 2
    	pfname = sprintf('../resources/pareto_fronts/%s.2D.pf', problem);
    else
        pfname = sprintf('../resources/pareto_fronts/%s.pf', problem);
    end
    
    
    %%     
    if strcmp(problem,'UF1')==1 || strcmp(problem,'UF2')==1 || strcmp(problem,'UF3')==1 || strcmp(problem,'UF4')==1 || ...
        strcmp(problem,'UF5')==1 || strcmp(problem,'UF6')==1 || strcmp(problem,'UF7')==1 || strcmp(problem,'ZDT1')==1 || ...
         strcmp(problem,'ZDT2')==1 ||  strcmp(problem,'ZDT3')==1 ||  strcmp(problem,'ZDT4')==1 || strcmp(problem,'ZDT6')==1 || ...
         strcmp(problem,'LZ09_F1')==1 || strcmp(problem,'LZ09_F2')==1 || strcmp(problem,'LZ09_F3')==1 || strcmp(problem,'LZ09_F4')==1 || ...
         strcmp(problem,'LZ09_F5')==1 || strcmp(problem,'LZ09_F7')==1 || strcmp(problem,'LZ09_F8')==1 || strcmp(problem,'LZ09_F9')==1

         points = 100;
    elseif strcmp(problem,'UF8')==1 || strcmp(problem,'UF9')==1 || strcmp(problem,'UF10')==1 || strcmp(problem,'LZ09_F6')==1 || ...
        strcmp(problem,'DTLZ1')==1 || strcmp(problem,'DTLZ2')==1 || strcmp(problem,'DTLZ3')==1 || strcmp(problem,'DTLZ4')==1 || ...
        strcmp(problem,'DTLZ5')==1 || strcmp(problem,'DTLZ6')==1 || strcmp(problem,'DTLZ7')==1 

        points = 150;
    else
    	points = 100;
    end
    
end